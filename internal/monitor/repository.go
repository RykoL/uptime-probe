package monitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

type Repository interface {
	GetMonitors(ctx context.Context) ([]*Monitor, error)
	SaveMonitor(ctx context.Context, monitor *Monitor) error
}

type PostgresRepository struct {
	conn *pgxpool.Pool
	log  *slog.Logger
}

func NewRepository(conn *pgxpool.Pool, logger *slog.Logger) PostgresRepository {
	return PostgresRepository{conn, logger}
}

type monitorRecord struct {
	Id         int
	Name       string
	Interval   time.Duration
	Definition string
}

func (r *PostgresRepository) SaveMonitor(ctx context.Context, monitor *Monitor) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if pErr := tx.Rollback(ctx); pErr != nil && !errors.Is(pErr, pgx.ErrTxClosed) {
			r.log.Error("failed to rollback transaction", "error", pErr)
		}
	}()

	monitorQuery := `
		INSERT INTO uptime.monitor(interval, name) VALUES ($1, $2) RETURNING id;
	`

	var monitorId int
	err = tx.QueryRow(ctx, monitorQuery, monitor.Interval, monitor.Name).Scan(&monitorId)

	if err != nil {
		return fmt.Errorf("failed to insert monitor: %w", err)
	}

	probeQuery := `
		INSERT INTO uptime.probe(definition, monitor_id) VALUES ($1, $2);
	`
	_, err = tx.Exec(ctx, probeQuery, "{}", monitorId)
	if err != nil {
		return fmt.Errorf("failed to insert probe: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *PostgresRepository) GetMonitors(ctx context.Context) ([]*Monitor, error) {
	rows, err := r.conn.Query(ctx, `
		SELECT 
			monitor.id, name, interval, definition
		FROM uptime.monitor 
		JOIN uptime.probe ON uptime.monitor.id = uptime.probe.monitor_id`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []*Monitor
	for rows.Next() {
		var res monitorRecord
		err = rows.Scan(&res.Id, &res.Name, &res.Interval, &res.Definition)

		if err != nil {
			r.log.Error("Failed to scan row")
		}

		m, err := NewMonitorFromRecord(res)

		if err != nil {
			r.log.Error("Failed to map into a monitor")
		}

		results = append(results, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

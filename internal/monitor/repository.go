package monitor

import (
	"context"
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

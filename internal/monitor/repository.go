package monitor

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

type Repository struct {
	conn *pgxpool.Pool
	log  *slog.Logger
}

func NewRepository(conn *pgxpool.Pool, logger *slog.Logger) Repository {
	return Repository{conn, logger}
}

type monitorRecord struct {
	Id         int
	Name       string
	Interval   time.Duration
	Definition string
}

func (r *Repository) GetMonitors(ctx context.Context) ([]monitorRecord, error) {
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

	var results []monitorRecord
	for rows.Next() {
		var res monitorRecord
		err = rows.Scan(&res.Id, &res.Name, &res.Interval, &res.Definition)

		if err != nil {
			r.log.Error("Failed to scan row")
		}

		results = append(results, res)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

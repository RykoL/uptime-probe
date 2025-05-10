package web

import (
	"context"
	"github.com/RykoL/uptime-probe/web/model"
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func getMonitors(ctx context.Context, conn *pgxpool.Pool) ([]*model.Monitor, error) {
	rows, err := conn.Query(ctx, "SELECT name FROM uptime.monitor")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []*model.Monitor
	for rows.Next() {
		var res model.Monitor
		err = rows.Scan(&res.Name)

		if err != nil {
			return nil, err
		}

		res.Status = "Up"
		results = append(results, &res)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

type StatusPage struct {
	conn *pgxpool.Pool
}

func NewStatusPage(conn *pgxpool.Pool) StatusPage {
	return StatusPage{conn}
}

func (s *StatusPage) Monitors(w http.ResponseWriter, r *http.Request) {
	monitors, err := getMonitors(r.Context(), s.conn)
	if err != nil {
		return
	}

	static.Index(monitors).Render(r.Context(), w)
}

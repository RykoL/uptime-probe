package web

import (
	"context"
	"github.com/RykoL/uptime-probe/web/model"
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"sort"
	"time"
)

func getMonitors(ctx context.Context, conn *pgxpool.Pool) ([]*model.Monitor, error) {
	rows, err := conn.Query(ctx, `
	WITH AggregatedHeartbeats AS (
	    SELECT
        	m.id AS monitor_id,
			m.name as monitor_name,
        	time_bucket(m.interval, h.timestamp) AS bucket_start,
        	h.success as is_up
    	FROM uptime.heartbeat h
             JOIN uptime.monitor m ON h.monitor_id = m.id
    	GROUP BY m.id, bucket_start, is_up
	),
    RankedHeartbeats AS (
    	SELECT
        	monitor_id,
			monitor_name,
            bucket_start,
            is_up,
            ROW_NUMBER() OVER (PARTITION BY monitor_id ORDER BY bucket_start DESC) AS rn
        FROM AggregatedHeartbeats
     )
	SELECT
		monitor_id,
		monitor_name,
    	bucket_start,
    	is_up
	FROM RankedHeartbeats
	WHERE rn <= 31
	ORDER BY monitor_id, bucket_start ASC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	monitors := make(map[string]*model.Monitor)

	for rows.Next() {
		var monitorId int
		var monitorName string
		var bucketStart time.Time
		var isUp bool

		err := rows.Scan(&monitorId, &monitorName, &bucketStart, &isUp)
		if err != nil {
			return nil, err
		}

		probeResult := model.ProbeResult{
			Timestamp: bucketStart,
			Success:   isUp,
		}

		if m, ok := monitors[monitorName]; ok {
			m.Results = append(m.Results, probeResult)
		} else {
			monitors[monitorName] = &model.Monitor{
				Id:      monitorId,
				Name:    monitorName,
				Results: []model.ProbeResult{probeResult},
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var results []*model.Monitor
	for _, m := range monitors {
		results = append(results, m)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Id < results[j].Id
	})

	return results, nil
}

type StatusPage struct {
	conn   *pgxpool.Pool
	logger *slog.Logger
}

func NewStatusPage(conn *pgxpool.Pool, logger *slog.Logger) StatusPage {
	return StatusPage{conn, logger}
}

func (s *StatusPage) Monitors(w http.ResponseWriter, r *http.Request) {
	monitors, err := getMonitors(r.Context(), s.conn)
	if err != nil {
		s.logger.Error("Failed retrieving monitors", "error", err)
		monitors = make([]*model.Monitor, 0)
	}

	static.Index(monitors).Render(r.Context(), w)
}

package main

import (
	"context"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/db"
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/RykoL/uptime-probe/web"
	"github.com/RykoL/uptime-probe/web/static"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if len(os.Args) < 2 {
		logger.Error("Missing required parameter config")
		os.Exit(1)

	}

	location, _ := time.LoadLocation("UTC")
	time.Local = location

	cfg, err := config.LoadFromFile(os.Args[1])

	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	dbpool, err := db.CreateDBPool(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	defer dbpool.Close()

	repository := monitor.NewRepository(dbpool, logger)

	manager := monitor.NewManager(logger, &repository)
	err = manager.Initialize(context.Background(), cfg)
	if err != nil {
		log.Fatal("Failed to initialize MonitorManager", err)
	}

	ctx, _ := context.WithCancel(context.Background())
	err = manager.Run(ctx)
	defer manager.Stop()

	if err != nil {
		panic(err)
	}

	statusPage := web.NewStatusPage(dbpool, logger)

	http.Handle("GET /static/", http.StripPrefix("/static/", static.NewEmbeddedFileHandler(static.Assets, http.NotFoundHandler())))
	http.HandleFunc("GET /", statusPage.Monitors)
	logger.Info("Starting web server on port :8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/db"
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/RykoL/uptime-probe/web/model"
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/a-h/templ"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if len(os.Args) < 2 {
		logger.Error("Missing required parameter config")
		os.Exit(1)
	}
	cfg, err := config.LoadFromFile(os.Args[1])

	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	dbpool, err := db.CreateDBPool(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	defer dbpool.Close()

	manager := monitor.NewManager(logger, nil)
	manager.ApplyConfig(cfg)
	manager.Run()

	monitors := []model.Monitor{
		{Name: "GCP", Status: "Up"},
		{Name: "AWS", Status: "Up"},
		{Name: "Azure", Status: "Unknown"},
		{Name: "StackIt", Status: "Pending"},
		{Name: "IONOS", Status: "Down"},
	}
	http.Handle("/", templ.Handler(static.Index(monitors)))

	http.ListenAndServe(":8080", nil)
}

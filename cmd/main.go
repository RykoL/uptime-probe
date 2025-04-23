package main

import (
	"github.com/RykoL/uptime-probe/config"
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

	manager := monitor.NewManager(logger)
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

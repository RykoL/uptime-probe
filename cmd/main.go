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
	cfg, err := config.LoadFromFile(os.Args[1])

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	manager := monitor.NewManager(logger)
	manager.ApplyConfig(cfg)

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

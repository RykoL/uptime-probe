package main

import (
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	m := monitor.NewMonitor("My Monitor")
	http.Handle("/", templ.Handler(static.Index([]monitor.Monitor{m})))

	http.ListenAndServe(":8080", nil)
}

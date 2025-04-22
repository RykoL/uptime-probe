package main

import (
	"github.com/RykoL/uptime-probe/web/model"
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
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

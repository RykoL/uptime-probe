package main

import (
	"github.com/RykoL/uptime-probe/web/static"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(static.Index()))

	http.ListenAndServe(":8080", nil)
}

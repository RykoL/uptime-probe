package static

import (
	"embed"
	"net/http"
)

type EmbeddedFileHandler struct {
	fs   http.FileSystem
	next http.Handler
}

func (h EmbeddedFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := h.fs.Open(r.URL.Path); err == nil {
		// File exists in embedded FS, serve it
		http.FileServer(h.fs).ServeHTTP(w, r)
		return
	}
	// File not found in embedded FS, pass to the next handler
	h.next.ServeHTTP(w, r)
}

func NewEmbeddedFileHandler(fs embed.FS, next http.Handler) EmbeddedFileHandler {
	return EmbeddedFileHandler{
		fs:   http.FS(fs), // Convert embed.FS to http.FileSystem
		next: next,
	}
}

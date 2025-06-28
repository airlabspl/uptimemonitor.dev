package handler

import (
	"net/http"
	"os"
	"selfhosted/ui"
	"strings"
)

func UI(w http.ResponseWriter, r *http.Request) {
	f, err := ui.FS().Open(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil && os.IsNotExist(err) {
		r.URL.Path = "/"
	}

	if err == nil {
		defer f.Close()
	}

	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
	} else {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Header().Set("Pragma", "public")
		w.Header().Set("Expires", "max-age=31536000")
	}

	http.FileServerFS(ui.FS()).ServeHTTP(w, r)
}

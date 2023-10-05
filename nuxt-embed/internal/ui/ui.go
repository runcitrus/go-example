package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed web/.output/public/_nuxt/* web/.output/public/favicon.ico
	public embed.FS

	//go:embed web/.output/public/index.html
	indexHTML []byte
)

func webSPA() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})
}

func webStatic() http.HandlerFunc {
	fsPublic, _ := fs.Sub(public, "web/.output/public")
	fileServer := http.FileServer(http.FS(fsPublic))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		fileServer.ServeHTTP(w, r)
	})
}

func WebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/_nuxt/", webStatic())
	mux.HandleFunc("/favicon.ico", webStatic())
	mux.HandleFunc("/", webSPA())
	return mux
}

package static

import (
	"embed"
	"net/http"

	"webimizer.dev/webimizer"
)

//go:embed favicon.ico favicon-16x16.png favicon-32x32.png apple-touch-icon.png robots.txt
var static embed.FS

func ServeStaticFiles() {
	http.Handle("/favicon.ico",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(static)).ServeHTTP(rw, r)
		}))
	http.Handle("/favicon-16x16.png",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(static)).ServeHTTP(rw, r)
		}))
	http.Handle("/favicon-32x32.png",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(static)).ServeHTTP(rw, r)
		}))
	http.Handle("/apple-touch-icon.png",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(static)).ServeHTTP(rw, r)
		}))
	http.Handle("/robots.txt",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(static)).ServeHTTP(rw, r)
		}))
}

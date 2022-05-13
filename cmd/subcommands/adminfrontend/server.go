package adminfrontend

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"webimizer.dev/webimizer"
)

//go:embed assets
var content embed.FS

func (p *adminFrontendCmd) runServer() {
	webimizer.DefaultHTTPHeaders = [][]string{
		{"x-content-type-options", "nosniff"},
		{"x-frame-options", "SAMEORIGIN"},
		{"x-xss-protection", "1; mode=block"},
	} // define web application default HTTP response headers
	http.Handle("/assets/",
		webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			http.FileServer(http.FS(content)).ServeHTTP(rw, r)
		})) // serve web static assets

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", p.host, p.port), nil)) // Start server
}

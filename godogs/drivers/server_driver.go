package drivers

import (
	"fmt"
	"html"
	"net/http"
)

var server *http.Server

const addr = "127.0.0.1:3001"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Application Up, %q", html.EscapeString(r.URL.Path))
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/health", healthHandler)
	return r
}

// CloseServer function
func CloseServer() {
	server.Close()
}

// RunApp function
func RunApp() {
	server = &http.Server{
		Addr:    addr,
		Handler: handler(),
	}
	go server.ListenAndServe()
}

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	handler := &MyHandler{}

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		})
		http.ListenAndServe(":80", nil)
	}()

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}

	http2.ConfigureServer(&server, &http2.Server{})

	server.ListenAndServeTLS("cert.pem", "key.pem")
}

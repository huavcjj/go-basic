package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	// h := r.Header
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}

func main() {
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}

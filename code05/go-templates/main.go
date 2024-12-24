package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tmpl := template.Must(template.ParseFiles("home.html", "page_header.html", "page_body.html"))
		tmpl := template.Must(template.ParseGlob("templates/*.html"))

		err := tmpl.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", nil)
}

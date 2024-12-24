package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tmpl := template.Must(template.ParseFiles("home.html", "page_header.html", "page_body.html"))
		tmpl := template.Must(template.ParseGlob("templates/*.html"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
			Features    []string
		}{
			Name:        "Golang",
			Title:       "Go テンプレートの例",
			Description: "これは、Goのテンプレートを使用したシンプルな例です。テンプレートは別ファイルで定義され、template.Must 関数を使用して解析されます。データは構造体を通じてテンプレートに渡されます。この構造体を使用してテンプレートをレンダリングします。",
			Socials: map[string]string{
				"Twitter":  "https://twitter.com/golang",
				"Facebook": "https://www.facebook.com/golang",
				"GitHub":   "https://github.com/golang",
			},
			Features: []string{
				"テンプレートの解析",
				"テンプレートのレンダリング",
			},
		}

		err := tmpl.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", nil)
}

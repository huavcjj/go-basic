package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	// 	tmpl := `
	// <!DOCTYPE html>
	// <html>
	//   <head>
	// 	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	// 	<title>Go Web Programming</title>
	//   </head>
	//   <body>
	// 	{{ . }}
	//   </body>
	// </html>`
	// 	t, _ := template.New("tmpl.html").Parse(tmpl)

	// t, _ := template.ParseGlob("*.html")

	// t, _ := template.ParseFiles("tmpl.html")

	//エラーを処理する場合
	// t := template.Must(template.ParseFiles("tmpl.html"))
	//関数Mustは、テンプレートへのポインタを返す関数とエラーをラップして、エラーがnilでない場合はパニックを発生させる。
	// t.Execute(w, "Hello World!")

	//セットのテンプレートを実行する場合
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
	t.ExecuteTemplate(w, "t2.html", "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

// % curl -i 127.0.0.1:8080/process

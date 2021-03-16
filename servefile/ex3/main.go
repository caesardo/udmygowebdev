package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

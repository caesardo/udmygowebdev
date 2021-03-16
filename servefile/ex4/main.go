package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

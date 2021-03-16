package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func index(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("Template failed to execute ", err)
	}
}

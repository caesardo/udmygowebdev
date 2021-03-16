package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml", "me.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Headers       http.Header
		ContentLength int64
		Form          map[string][]string
	}{
		req.Method,
		req.URL,
		req.Header,
		req.ContentLength,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Dawg")
}

func me(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	name := req.Form["name"]
	tpl.ExecuteTemplate(w, "me.gohtml", name)
}

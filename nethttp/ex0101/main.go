package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")

	// // buffered writer
	// writer := bufio.NewWriter(w)
	// writer.WriteString("Hello World")
	// writer.Flush()
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Dawg")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi My name is Novaldo")
}

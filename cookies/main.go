package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "counter-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	current, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	current++
	c.Value = strconv.Itoa(current)

	http.SetCookie(w, c)

	fmt.Fprintln(w, "COOKIE WRITTEN")
	fmt.Fprintln(w, "count :", c.Value)

}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counter-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "web count :", c)
}

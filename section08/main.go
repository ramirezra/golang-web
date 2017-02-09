package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count int

func main() {
	http.HandleFunc("/", counter)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	count++
	http.SetCookie(w, &http.Cookie{
		Name:  "Counter",
		Value: strconv.Itoa(count),
	})

}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Counter")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sally sells sea shells by the sea shore!")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

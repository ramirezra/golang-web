package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Robinson-Key", "This is from my program.")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h4>Robinson's Test Page</h4>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

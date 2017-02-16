package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

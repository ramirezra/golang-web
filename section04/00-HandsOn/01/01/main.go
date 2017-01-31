// ListenAndServe on port ":8080" using the default ServeMux.
//
// Use HandleFunc to add the following routes to the default ServeMux:
//
// "/"
// "/dog/"
// "/me/
//
// Add a func for each of the routes.
//
// Have the "/me/" route print out your name.

package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "me.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

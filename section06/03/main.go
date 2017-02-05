package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dogs)

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))

}

func dogs(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

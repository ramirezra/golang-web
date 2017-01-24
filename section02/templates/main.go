package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer newfile.Close()

	err = tpl.Execute(newfile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

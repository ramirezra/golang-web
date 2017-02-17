package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	if r.Method == http.MethodPost {
		// filename := r.FormValue("filename")
		file, header, err := r.FormFile("filename")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		// create new file
		workdirectory, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(workdirectory, "public", "pics", header.Filename)
		newfile, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer newfile.Close()
		// copy file
		file.Seek(0, 0)
		io.Copy(newfile, file)

		if !strings.Contains(c.Value, header.Filename) {
			c.Value += "|" + header.Filename
		}
	}
	http.SetCookie(w, c)

	return c
}

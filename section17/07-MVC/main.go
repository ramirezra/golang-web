package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ramirezra/golang-web/section17/07-MVC/books"
)

func main() {
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/books", books.BooksIndex)
	http.HandleFunc("/books/show", books.BooksShow)
	http.HandleFunc("/books/create", books.BooksCreateForm)
	http.HandleFunc("/books/create/process", books.BooksCreateProcess)
	http.HandleFunc("/books/update", books.BooksUpdateForm)
	http.HandleFunc("/books/update/process", books.BooksUpdateProcess)
	http.HandleFunc("/books/delete/process", books.BooksDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

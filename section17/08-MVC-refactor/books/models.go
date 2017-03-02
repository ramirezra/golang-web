package books

import (
	"net/http"

	"github.com/ramirezra/golang-web/section17/08-MVC-refactor/config"
)

// Book exported
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks exported
func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
}

func OneBook() {

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
}

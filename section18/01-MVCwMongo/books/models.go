package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ramirezra/golang-web/section18/01-MVCwMongo/config"

	"gopkg.in/mgo.v2/bson"
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
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}

// OneBook exported to handlers.go
func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. BadRequest")
	}

	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// PutBook exported to handlers.go
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. BadRequest")
	}

	// convert form values
	float, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("400. BadRequest")
	}
	bk.Price = float32(float)

	// insert values
	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())

	}
	return bk, nil
}

// UpdateBook exported to handlers.go
func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("406. Not Acceptable. Fields can't be blank")
	}

	// convert form values
	float, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(float)

	// update values
	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// DeleteBook exported to handlers.go
func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	// delete bookstore
	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

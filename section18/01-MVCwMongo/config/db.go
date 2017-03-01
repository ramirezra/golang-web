package config

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// DB exported
var DB *mgo.Database

// Books collection exported
var Books *mgo.Collection

func init() {
	// var err error
	// DB, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You connected to your database.")

}

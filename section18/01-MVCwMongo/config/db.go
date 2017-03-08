package config

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// DB exported
var DB *mgo.Database

// Books exported
var Books *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://localhost/bookstore")
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

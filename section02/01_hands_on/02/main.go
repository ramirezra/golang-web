// 1. Create a data structure to pass to a template which
// * contains information about California hotels including Name, Address, City, Zip, Region
// * region can be: Southern, Central, Northern
// * can hold an unlimited number of hotels

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}
type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Holiday Inn",
			Address: "101 Dalmation Drive",
			City:    "London",
			Zip:     "12345",
			Region:  "Southern",
		},
		hotel{
			Name:    "Budget Inn",
			Address: "2 Saint Patrick Way",
			City:    "Pittsburgh",
			Zip:     "15233",
			Region:  "Central",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}

// * Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Post is Eported type
type Post struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	AdjClose float64
}

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

func stocks(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	// var table []string

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1

	record, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var posts []Post

	for i, item := range record {
		open, _ := strconv.ParseFloat(item[1], 64)
		high, _ := strconv.ParseFloat(item[2], 64)
		low, _ := strconv.ParseFloat(item[3], 64)
		close, _ := strconv.ParseFloat(item[4], 64)
		volume, _ := strconv.ParseInt(item[5], 0, 0)

		if i == 0 {
			continue
		}

		post := Post{Date: item[0], Open: open, High: high, Low: low, Close: close, Volume: volume}
		posts = append(posts, post)
	}

	error := tpl.ExecuteTemplate(w, "stocks.gohtml", posts)
	if error != nil {
		log.Fatalln(error)
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
	http.HandleFunc("/stocks", stocks)

	http.ListenAndServe(":8080", nil)
}

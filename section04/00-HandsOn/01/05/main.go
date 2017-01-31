// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc
//
// Contstraint: Do not change anything outside of func main
//
// Hints:
//
// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```
//
// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```
//
// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//   		mux.Handle(pattern, HandlerFunc(handler))
//   }
// ```

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
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.HandleFunc("/stocks", http.HandlerFunc(stocks))

	http.ListenAndServe(":8080", nil)
}

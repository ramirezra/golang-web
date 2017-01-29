package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

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

func main() {
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

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer newfile.Close()

	error := tpl.Execute(newfile, posts)
	if err != nil {
		log.Fatalln(error)
	}
}

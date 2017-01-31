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
	Volume   float64
	AdjClose float64
}

//
// func makeRecord(row []string) record {
// 	open, _ := strconv.ParseFloat(row[1], 64)
// 	high, _ := strconv.ParseFloat(row[2], 64)
// 	low, _ := strconv.ParseFloat(row[3], 64)
// 	close, _ := strconv.ParseFloat(row[4], 64)
// 	volume, _ := strconv.ParseFloat(row[5], 64)
// 	// adjusted, _ := strconv.ParseFloat(row[6], 64)
//
// 	return record{
// 		Date:   row[0],
// 		Open:   open,
// 		High:   high,
// 		Low:    low,
// 		Close:  close,
// 		Volume: volume,
// 		// Adjusted: adjusted,
// 	}
// }

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

	for _, item := range record {
		open, _ := strconv.ParseFloat(item[1], 64)
		// high, _ := strconv.ParseFloat(record[2], 64)
		// low, _ := strconv.ParseFloat(record[3], 64)
		// close, _ := strconv.ParseFloat(record[4], 64)
		// volume, _ := strconv.ParseFloat(record[5], 64)

		post := Post{Date: item[0], Open: open}
		posts = append(posts, post)
	}

	error := tpl.Execute(os.Stdout, posts)
	if err != nil {
		log.Fatalln(error)
	}
}

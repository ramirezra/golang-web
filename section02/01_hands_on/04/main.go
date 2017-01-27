package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type record struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

type records []record

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	high, _ := strconv.ParseFloat(row[2], 64)
	low, _ := strconv.ParseFloat(row[3], 64)
	close, _ := strconv.ParseFloat(row[4], 64)
	volume, _ := strconv.ParseFloat(row[5], 64)
	// adjusted, _ := strconv.ParseFloat(row[6], 64)

	return record{
		Date:   row[0],
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
		// Adjusted: adjusted,
	}
}

func main() {
	data, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer data.Close()

	// var table []string

	for {
		record, err := csv.NewReader(data).ReadAll()
		if err != nil {
			log.Fatalln(err)
		}
		for i, row := range record {
			if i == 0 {
				continue
			}
			fmt.Println(i, row)

			// record := makeRecord(row)
		}
	}

	// err := tpl.Execute(os.Stdout, record)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}

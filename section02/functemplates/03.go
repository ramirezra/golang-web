package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl3 *template.Template

func init() {
	tpl3 = template.Must(template.New("").Funcs(fm).ParseFiles("tpl03.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func main() {
	err := tpl3.ExecuteTemplate(os.Stdout, "tpl03.gohtml", 5)
	if err != nil {
		log.Fatalln(err)
	}
}

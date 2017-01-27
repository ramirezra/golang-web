// 1. Create a data structure to pass to a template which
// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	ItemName, Price string
}

type meal struct {
	MealType string
	Items    []item
}

type restaurant struct {
	ResName string
	Meal    []meal
}

type restaurants struct {
	Restaurants []restaurant
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main() {

	restaurants := []restaurant{
		restaurant{
			"Beni Hana",
			[]meal{
				meal{
					MealType: "Lunch",
					Items: []item{
						item{"Egg Roll", "$4.95"},
						item{"Udon Noodles", "$7.95"},
					},
				},
				meal{
					MealType: "Dinner",
					Items: []item{
						item{"Sushi", "$18.95"},
						item{"Edamame", "$12.95"},
					},
				},
			},
		},
		restaurant{
			"International Pancake House",
			[]meal{
				meal{
					MealType: "Breakfast",
					Items: []item{
						item{"Silver Dollar Pancakes", "$4.95"},
						item{"Wester Omelet", "$7.95"},
					},
				},
				meal{
					MealType: "Lunch",
					Items: []item{
						item{"Hot Dog", "$4.95"},
						item{"Pizza", "$7.95"},
					},
				},
				meal{
					MealType: "Dinner",
					Items: []item{
						item{"Spaghetti", "$14.95"},
						item{"Steak and Eggs", "$17.95"},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(restaurants)

}

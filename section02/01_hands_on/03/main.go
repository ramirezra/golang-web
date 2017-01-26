// 1. Create a data structure to pass to a template which
// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

package main

import "fmt"

type food struct {
	FoodName string
	Price    string
}

type foods []food

type meal struct {
	foods
	Meal string
}

type meals []meal

type restaurant struct {
	ResName string
	meals
}

func main() {
	n := restaurant{
		ResName: "Hoosiers",
		meals: []meal{
			Meal: "Breakfast",
		},
	}

	fmt.Println(n)
}

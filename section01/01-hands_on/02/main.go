package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak() string
}

func (p person) speak() {
	fmt.Println(p.firstname)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.licenseToKill)
}

func main() {
	p1 := person{
		"Robinson",
		"Ramirez",
	}

	s2 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	p1.speak()
	s2.speak()
}

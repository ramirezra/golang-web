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

func (p person) speak() {
	fmt.Printf("Hello %s, you have a package.\n", p.firstname)
}

func (sa secretAgent) speak() {
	fmt.Printf("Hello %s, you have a new mission.\n", sa.firstname)
}

type human interface {
	speak()
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
	saysomething(p1)
	saysomething(s2)
}

func saysomething(h human) {
	h.speak()
}

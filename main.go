package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)
	fmt.Printf("%d %d\n", xi[0], xi[4])

	m := map[string]int{
		"Todd":     45,
		"Job":      42,
		"Robinson": 36,
	}
	fmt.Println(m)
	fmt.Println(m["Robinson"])

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()

	saySomething(p1)
	saySomething(sa1)
	saySomething(sa1.person)
}

package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

// Empty function. Setting up for Exercise 04.
func main() {
	p1 := person{"Robinson", "Ramirez", []string{"pizza", "oranges", "lasagna"}}

	fmt.Println(p1.favFood)

	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}

	s := p1.walk()
	fmt.Println(s)
}

func (p person) walk() string {
	return fmt.Sprintf("%s is walking.", p.fname)
}

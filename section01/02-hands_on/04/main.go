package main

import "fmt"

type person struct {
	fname string
	lname string
}

// Empty function. Setting up for Exercise 04.
func main() {
	p1 := person{"Robinson", "Ramirez"}

	fmt.Println(p1.fname)
}

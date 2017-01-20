package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

var g1 gator
var x int
var f flamingo

func main() {
	g1 = 3
	fmt.Printf("%v: %T\n", g1, g1)
	fmt.Printf("%d: %T\n", x, x)
	x = int(g1)
	fmt.Printf("%d: %T\n", x, x)
	g1.greeting()
	bayou(g1)
	bayou(f)
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(s swampCreature) {
	s.greeting()
}

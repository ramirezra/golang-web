package main

import "fmt"

type gator int

var g1 gator
var x int

func main() {
	g1 = 3
	fmt.Printf("%v: %T\n", g1, g1)
	fmt.Printf("%d: %T\n", x, x)
	x = int(g1)
	fmt.Printf("%d: %T\n", x, x)
	g1.greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

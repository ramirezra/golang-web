package main

import "fmt"

type gator int

var g1 gator

func main() {
	g1 = 3
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
}

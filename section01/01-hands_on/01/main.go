package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

func (c circle) area() float64 {
	return c.radius * math.Pi
}

func main() {
	s := rectangle{6, 6}
	c := circle{2.0}
	// fmt.Printf("The area of a square is %f.\n", rectangle.area(s))
	// fmt.Printf("The area of the circle is %f.\n", circle.area(c))
	info(s)
	info(c)
}

func info(s shape) {
	fmt.Printf("The area of the shape %T is %f.\n", s, s.area())
}

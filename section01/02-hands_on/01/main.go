package main

import "fmt"

func main() {
	x := []int{3, 4, 5, 6}

	for i, v := range x {
		// fmt.Println(i)
		fmt.Printf("%d: %d\n", i, v)
	}

}

package main

import "fmt"

func main() {
	m := map[string]int{
		"Robinson": 36,
		"Dave":     42,
		"Billy":    45,
	}

	fmt.Println(m)

	for s, _ := range m {
		fmt.Println(s)
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}

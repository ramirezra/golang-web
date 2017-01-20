package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{vehicle{2, "Red"}, true}
	s := sedan{vehicle{4, "Green"}, false}
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}

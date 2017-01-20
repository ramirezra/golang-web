// Give a method to both the “truck” and “sedan” types with the following signature
// 	transportationDevice() string
// Have each func return a string saying what they do. Create a value of type truck and populate the fields. Create a value of type sedan and populate the fields. Call the method for each value.

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

type transportation interface {
	transportationDevice()
}

func main() {
	t := truck{vehicle{2, "Red"}, true}
	s := sedan{vehicle{4, "Green"}, false}
	fmt.Println(t.doors)
	fmt.Println(s.doors)

	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())

	fmt.Println(report(s))
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("This is a truck for weekend warriors.")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("This is a nice luxury car.")
}

func report(v transportation) {
	v.transportationDevice()
}

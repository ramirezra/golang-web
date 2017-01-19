package main

type Person struct {
	Firstname string
	Lastname  string
}

type SecretAgent struct {
	Person
	LicenseToKill bool
}

type human interface {
	speak() string
}

func (p Person) speak() string {
	return p.Firstname
}

func (sa SecretAgent) speak() string {
	return sa.Person.Firstname
}

func main() {
	p1 := Person{
		"Robinson",
		"Ramirez",
	}

	s2 := SecretAgent{
		Person{
			"James",
			"Bond",
		},
		true,
	}
	says(p1)
	says(s2)
}

func says(h human) {
	h.speak()
}

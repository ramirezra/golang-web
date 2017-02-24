package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Id        string
	Firstname string
	Lastname  string
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	file, err := os.Create("data.json")
	check(err)
	defer file.Close()

	u := user{"1", "Robinson", "Ramirez"}
	// u := []byte{115, 111, 109, 101}
	json, err := json.Marshal(u)
	check(err)
	file.Write(json)
	check(err)
	fmt.Println("Save data.json with:", file)
}

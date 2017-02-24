package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// Id was of type string before

// StoreUsers exported
func StoreUsers(m map[string]User) {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(m)
}

// LoadUsers exported
func LoadUsers() map[string]User {
	m := make(map[string]User)

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

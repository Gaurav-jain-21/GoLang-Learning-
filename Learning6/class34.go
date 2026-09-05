package main

import (
	"encoding/json"
	"fmt"
)

type User3 struct {
	Name    string
	Age     int
	Email   string
	Hobbies []string
}

func class34() {
	user := User3{
		Name:    "Gaurav",
		Age:     22,
		Email:   "gjain@gmail.com",
		Hobbies: []string{"coding", "gaming", "reading"},
	}

	jsonData, err := json.Marshal(user)

	if err!= nil{
		fmt.Println("Error",err)
		return 
	}
	fmt.Println("Json (bytes):",jsonData)
	fmt.Println("Json (string):",string(jsonData))
}
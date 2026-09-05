package main

import (
	"encoding/json"
	"fmt"
)

type User5 struct {
	Name  string
	Age   int
	Email string
}

func class36() {
	jsonString := `{"name":"Gaurav","age":22,"email":"gjain@gmail.com"}`

	var user User5
	err := json.Unmarshal([]byte(jsonString),&user)

	if err!= nil{
		fmt.Println("Error",err)
		return 
	}
	fmt.Println("Name:",user.Name)
	fmt.Println("Age:",user.Age)
	fmt.Println("Email:",user.Email)

}
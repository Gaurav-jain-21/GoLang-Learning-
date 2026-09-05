package main

import (
	"encoding/json"
	"fmt"
)

type User4 struct {
	Name  string
	Email string
}

func class35() {
	user := User4{Name: "Gaurav", Email: "gjain@gmail.com"}
	jsonData, _ := json.MarshalIndent(user,""," ")
	fmt.Println(string(jsonData))

}

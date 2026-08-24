package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	user := User{"Gaurav","gjain7764@gmail.com",true,16}
	fmt.Println(user)
}

type User struct {
	Name string 
	Email string
	Status bool
	Age int 
}


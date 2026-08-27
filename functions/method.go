package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	hitesh := User{"Hitesh", "hti@gmail.com", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("name is %v and email is %v", hitesh.Name, hitesh.Email)
	fmt.Println("")
	hitesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ",u.Status)
}
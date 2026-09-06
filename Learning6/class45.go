package main

import "fmt"

type User6 struct {
	Name  string
	Email string
	Age   int
}

func (u User6) String() string {
	return fmt.Sprintf("User{Name: %s, Email: %s , Age: %d}", u.Name, u.Email, u.Age)
}

func class45(){
	user:=User6{"Garuav","gjain@gmail",22}
	fmt.Println(user)
}
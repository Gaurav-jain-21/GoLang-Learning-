package main

import "fmt"

type User1  struct{
	Name string
	Email string 
}
func updateEmail(u *User1, newEmail string){
	u.Email= newEmail
}

func class26(){
	user:= User1{Name: "Gaurav",Email: "oldEam"}
	fmt.Println("before", user.Email)
	updateEmail(&user,"new@gamilc.om")
	fmt.Println("after ",user.Email)
}
package main

import "fmt"

type User2 struct{
	Id int 
	Username string 
	Email string 
	IsActive bool
}

func class4(){
	newUser:= User2{
		Id: 101,
		Username: "Gaurav",
		Email: "gjain@gmail.com",
		IsActive:true ,
	}
	fmt.Printf("User %s has the email %s \n",newUser.Username, newUser.Email)
}
package main

import "fmt"

type User struct{
	Id int 
	Username string 
	Email string 
	IsActive bool
}

func class4(){
	newUser:= User{
		Id: 101,
		Username: "Gaurav",
		Email: "gjain@gmail.com",
		IsActive:true ,
	}
	fmt.Printf("User %s has the email %s \n",newUser.Username, newUser.Email)
}
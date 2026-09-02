package main

import "fmt"
func login(){
	const correctUsername= "admin"
	const correctPass="123123"
	var username, password string
	fmt.Println("Enter the username ")
	fmt.Scanln(&username)
	fmt.Println("Enter the password ")
	fmt.Scanln(&password)

	isUsernameCorrect:= username==correctUsername
	isPasswordCorrect:= password==correctPass
	if isUsernameCorrect && isPasswordCorrect{
		fmt.Println("Login Successfull Welcome")
	}else{
		fmt.Println("Invalid username or password")
	}
}
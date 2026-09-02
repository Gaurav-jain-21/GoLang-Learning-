package main

import "fmt"

func class2(){
	userAge:=16
	isAdmin:= true
	if  userAge<18 {
		fmt.Println("Error: user must be 18 or older")
		return 
	}
	if isAdmin!=true{
		fmt.Println("Error: Admin Privileges required ")
		return
	}else{
		fmt.Println("Success: Resource deleted")
	}
}
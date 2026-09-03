package main

import "fmt"
type UserID int 
type Email string 
type Status string 

func class22(){
	var id UserID = 101
	var email Email="gjain@gmail.com"
	var status Status="active"

	fmt.Printf("Id is %d and type is %T\n",id,id)
	fmt.Printf("Email is %s and type is %T\n", email, email)
	fmt.Printf("status is %s type is %T\n",status, status)
}
package main

import "fmt"
func sic(){
	var principal float64
	var rate float64
	var time float64
	fmt.Println("enter the principal amount: ")
	fmt.Scanln(&principal)
	fmt.Println("Enter the Rate of interest")
	fmt.Scanln(&rate)
	fmt.Println("Time to interest")
	fmt.Scanln(&time)

	interest:=(principal*rate*time)/100
	totalAmount:=principal+interest

	fmt.Println("Simple Interest : ", interest)
	fmt.Println("Total Amount: ",totalAmount)
}
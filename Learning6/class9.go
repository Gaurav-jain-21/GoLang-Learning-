package main

import "fmt"

func class9(){
	name:="Gaurav"
	age:= 22
	height:=1.73
	like:=true
	fmt.Printf("my name is %v and my age is %v with my height %v my love for go is %v", name, age, height, like)

	//(=====Swap Two Numbers======)
	num1:=8
	num2:=16
	fmt.Println("num1 is ",num1)
	fmt.Println("num2 is ", num2)
	temp:=num1
	num1= num2
	num2=temp
	fmt.Println("the swap value of num1 is ", num1)
	fmt.Println("the swap value of num2 is ",num2)

	//(===== Circle Calculator======)
	fmt.Println("Enter the radius of circle")
	var radius int
	fmt.Scanln(&radius)
	fmt.Println("Area of circle is ", 3.14159*(float64(radius)*float64(radius)))
	fmt.Println("Circumference of the circle is ",3.14159* 2*(float64(radius)))

}
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of Pointers")
	var ptr  *int 
	fmt.Println("value of pointer is ",ptr)
	myNumber := 23
	var ptr1 = &myNumber
	fmt.Println(ptr1)
	fmt.Println(*ptr1)

	
}
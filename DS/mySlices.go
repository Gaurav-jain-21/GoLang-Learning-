package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on the Slices")

	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitlist is %T\n",fruitList)
	fruitList = append(fruitList, "mango","banana")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:])
	
}
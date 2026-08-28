package main

import "fmt"
func chai6(){
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2]= "Potato"
	fruitList[3] = "pea"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList= [3]string{"potato","beans","carrot"}
	fmt.Println("Vegy list is : ", len(vegList))
}
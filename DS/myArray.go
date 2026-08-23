package main

import "fmt"
func main(){
	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string
	fruitList[0]="apple"
	fruitList[1]="tomato"
	fruitList[3]="peach"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println("vegy list is :",len(vegList))
}
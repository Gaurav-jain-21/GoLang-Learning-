package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of slices")
	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("type of the fruitList is %T \n", fruitList)
	fmt.Println("lenght of the slices",len(fruitList))
	fruitList = append(fruitList, "mango","banana")
	fmt.Println(fruitList)
	fruitList= append(fruitList[1:])
	fmt.Println(fruitList)
	highScore := make([]int, 4)
	highScore[0]= 234
	highScore[1]=76
	highScore[2]=85
	highScore[3]= 7
	// highScore[4]=3
	fmt.Println(highScore)
	highScore= append(highScore, 34353,632,9885)
	fmt.Println(highScore)
}


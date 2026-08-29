package main

import (
	"fmt"
	"math/rand"
	"time"
)
func chai10(){
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber:= rand.Intn(6)+1
	fmt.Println("Value of dice is ",diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
    case 3:
		fmt.Println("you can move to 3 spot")
	case 4:
		fmt.Println("you can move to 4 spot")
    default:
		fmt.Println("you can go to home")					

	}
}
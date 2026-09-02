package main

import "fmt"
func class13(){
	var secretNumber = 7
	var guess int 

	attempts:=0

	fmt.Println("Guess the number (1-10)")
	for guess != secretNumber{
		fmt.Println("your guess : ")
		fmt.Scanln(&guess)
		attempts++
		if guess< secretNumber{
			fmt.Println("Too low try again")
		}else{
			fmt.Println("Too high try again")
		}
	}
	fmt.Println("you guessed it")
}
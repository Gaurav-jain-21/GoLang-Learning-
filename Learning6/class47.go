package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func class47() {
	go printNumbers()  // Goroutine 1
	go printLetters()  // Goroutine 2
	
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}
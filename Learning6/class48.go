package main

import (
	"fmt"
	"time"
)

func class48() {
	fmt.Println("Starting tasks...")
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Background task:", i)
			time.Sleep(200 * time.Millisecond)

		}
	}()

	for i:=1; i<=3; i++{
		fmt.Println("Main task",i)
		time.Sleep(200*time.Millisecond)
	}
	time.Sleep(1*time.Second)
	fmt.Println("All done")
}
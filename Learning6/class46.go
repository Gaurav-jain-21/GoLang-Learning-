package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello from goroutine!", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func class46(){
	fmt.Println("Main started")
	go sayHello()

	fmt.Println("Main finished")
	time.Sleep(1*time.Second)
}
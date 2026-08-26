package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")
	days := []string{"Sunday","monday","tuesday","wednesday","friday","saturday"};
	fmt.Println(days)
	for d:=0; d<len(days); d++ {
		fmt.Println(days[d])
	}

}
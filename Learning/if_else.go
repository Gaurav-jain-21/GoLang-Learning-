package main

import "fmt"

func main() {
	age := 18

	if age > 15 {
		fmt.Println("the person is man")
	} else if age > 10 {
		fmt.Println("the person is boy")
	} else {
		fmt.Println("the perosn is boy")
	}
}
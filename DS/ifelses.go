package main

import "fmt"

func main() {
	fmt.Println("if else in golang")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "something else"
	}
	fmt.Println(result)
}
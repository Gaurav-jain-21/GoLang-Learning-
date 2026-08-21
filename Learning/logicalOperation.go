package main

import "fmt"

func main() {
	var role = "admin"
	var hasPermission = false
	if role == "admin" && hasPermission == true {
		fmt.Println("you can enjoy")
	} else {
		fmt.Println("you can't enjoy")
	}
}
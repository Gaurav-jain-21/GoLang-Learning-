package main

import "fmt"

const age=30

func main() {
	const name string = "golang"

	fmt.Println(age)
	fmt.Println(name)

	// name = "java"
	fmt.Println(name)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza : ");
	// comma ok | error ok
	input, _ := reader.ReadString('\n')
	fmt.Println(input);
}
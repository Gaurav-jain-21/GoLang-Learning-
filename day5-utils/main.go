package main

import (
	"day5-utils/strings"
	"fmt"
)

func main() {
	word := "radar"
	fmt.Printf("Is %s a palindrome ? %t\n", word, strings.IsPalindrome(word))

	reversed:= strings.Reverse("Gaurav")
	fmt.Println("Reversed : ",reversed)
	
}
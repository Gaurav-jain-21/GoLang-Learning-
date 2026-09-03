package main

import (
	"day5-utils/strings"
	"day5-utils/validators"
	"fmt"
)

func main() {
	word := "radar"
	fmt.Printf("Is %s a palindrome %t\n", word, strings.IsPalindrome(word))
	reversed := strings.Reverse("Gaurav")
	fmt.Println("Reversed:", reversed)
	
	// Using validators package
	email := "gaurav@example.com"
	fmt.Printf("Is '%s' valid? %t\n", email, validators.IsValidEmail(email))
	
	password := "12345"
	fmt.Printf("Is password strong? %t\n", validators.IsStrongPassword(password))
}
package main

import (
	"fmt"
	"strings"
)

func class31() {
	text := "  Hello, Gaurav!  "
	fmt.Println("Original: ", text)
	fmt.Println("Upper: ", strings.ToUpper(text))
	fmt.Println("Lower: ",strings.ToLower(text))
	fmt.Println("TrimSpace: ",strings.TrimSpace(text))
	fmt.Println("Contains Gaurav",strings.Contains(text,"Gaurav"))
	fmt.Println("Replaces: ",strings.Replace(text,"Gaurav","world",1))
	fmt.Println("Split: ",strings.Split("apple, banana, mango",","))
	filename:="document.pdf"
	fmt.Println("Has .pdf?",strings.HasSuffix(filename, ".pdf"))
	fmt.Println("Starts with doc?",strings.HasPrefix(filename,"doc"))
}
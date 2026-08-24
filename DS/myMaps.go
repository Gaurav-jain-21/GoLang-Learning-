package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	language:= make(map[string]string)
	language["js"]="JavaScript"
	language["py"]="Python"
	language["RB"]="Ruby"
	language["Java"]="java"

	fmt.Println(language)
	fmt.Println("js short for: ",language["js"])
	for key, value := range language{
		fmt.Printf("for key v, value is %v\n", key, value)

	}
}
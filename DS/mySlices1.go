package main

import "fmt"

func main() {
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int =2
	courses= append(courses,"golang")
	fmt.Println(courses)
	courses = append(courses[:index])
	fmt.Println(courses)
	courses= append(courses[:index],courses[index+1:]...)
}
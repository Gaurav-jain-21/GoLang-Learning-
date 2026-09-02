package main

import "fmt"
func main(){
	type Student struct{
		Name string
		Marks int
	}

	class:=[3]Student{
		{"Gaurav",85},
		{"Rahul",92},
		{"Priya",78},
	}
	for _,student:= range class{
		fmt.Println(student.Name,student.Marks)
	}
}
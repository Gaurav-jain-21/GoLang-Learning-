package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

type Employee1 struct {
	Person1
	EmployeeId int
	Department string
	Salary     float64
}

func class29() {
	emp := Employee1{
		Person1: Person1{
			"Gaurav", 22,
		},
		EmployeeId: 1001, Department: "Engineering", Salary: 75000,
	}
	fmt.Println("Name: ", emp.Name)
	fmt.Println("Age: ",emp.Age)
}
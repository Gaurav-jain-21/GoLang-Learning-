package main

import "fmt"

type Address struct {
	Street  string
	City    string
	Pincode int
}

type Employee struct {
	ID      int
	Name    string
	Salary  float64
	Address Address
}

func class21() {
	emp := Employee{
		ID:     101,
		Name:   "Gaurav",
		Salary: 75000.0,
		Address: Address{Street: "123 MG Road",
			City:    "Mumbai",
			Pincode: 400001,
		},
	}
	fmt.Println("Employee:", emp.Name)
	fmt.Println("City:", emp.Address.City) // Access nested field
	fmt.Println("Pincode:", emp.Address.Pincode)
}
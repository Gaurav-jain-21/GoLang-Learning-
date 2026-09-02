package main 

import "fmt"
type Person struct{
	Name string 
	Age int 
	Email string 
	City string 
}
func class19 (){
	p1:=Person{"Gaurav",22,"gjain@gmail.com","Mumbai"}

	p2:=Person{
		Name: "Rahul",
		Age:23,
		Email:"rahul@gmail.com",
		City: "Delhi",
	}
	var p3 Person
	p4:=Person{}
	p4.Name="Priya"
	p4.Age=23
	p4.Email="p@gmail.com"
	p4.City="Punea"

	fmt.Println("person1: ",p1)
	fmt.Println("person2: ",p2)
	fmt.Println("person3: ",p3)
	fmt.Println("person4: ",p4)
}
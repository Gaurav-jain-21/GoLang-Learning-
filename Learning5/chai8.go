package main

import "fmt"
func main(){
	fmt.Println("Structs in golang")
	// no inheritance in golang; NO super or parent
	gaurav := User{"Gaurav", "gjain@gmail.com",true, 22}
	fmt.Println(gaurav)
	fmt.Printf("Garuav detail are : %+v\n", gaurav)
	fmt.Printf("Name is %v and email is %v",gaurav.Name, gaurav.Email)
}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
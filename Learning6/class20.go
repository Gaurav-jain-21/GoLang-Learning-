package main 
import "fmt"

type Product struct{
	ID int 
	Name string
	Price float64
	Category string
	InStock bool
}
func class20(){
	laptop:=Product{
		ID: 101,
		Name: "mac book pro",
		Price: 124255.3,
		Category: "Electronics",
		InStock: true,
	}
	fmt.Println("Product name : ",laptop.Name)
	fmt.Println("Product price: ",laptop.Price)
	laptop.Price= 129999
	fmt.Println("Discount in the price ",laptop.Price)
}
package main

import "fmt"

type Product struct{
	ID int 
	Name string
	Price float64
	Qunatity int
}

func (p Product) Display(){
	fmt.Printf("ID: %d | %s | Rs%.2f | stock: %d\n",p.ID, p.Name, p.Price, p.Qunatity)
}

func (p *Product) Sell(quantity int){
	if quantity > p.Qunatity{
		fmt.Println("Not Enough stock")
		return 
	}
	p.Qunatity-=quantity
	total:= float64(quantity)*p.Price
	fmt.Printf("✅ Sold %d x %s for ₹%.2f\n", quantity, p.Name, total)
}
func (p *Product) Restock(quantity int){
	p.Qunatity+= quantity
    fmt.Printf("✅ Restocked %s by %d units\n", p.Name, quantity)
}

func main(){
	products:= []Product{
		{ID: 1, Name: "Laptop", Price: 50000, Qunatity: 10},
        {ID: 2, Name: "Mouse", Price: 500, Qunatity: 50},
        {ID: 3, Name: "Keyboard", Price: 1500, Qunatity: 30},
	}
	fmt.Println("====Product Catalog=====")
	for i:= range products{
		products[i].Display()
	}
	products[0].Sell(2)
}
package main

import "fmt"

func cart(){

	cart:=[]string{}
	var choice int 
	var item string
	for {
		fmt.Println("Welcome to the cart")
		fmt.Println("1. Add Item")
		fmt.Println("2. view Cart")
		fmt.Println("3. Exit")
		fmt.Println("Choice: ")
		fmt.Scanln(&choice)

		if choice ==1{
			fmt.Println("Enter the item name: ")
			fmt.Scanln(&item)
			cart=append(cart, item)
			fmt.Println("added ",item)
		}else if choice==2{
			fmt.Println("\n your cart: ")
			for i, product := range cart{
				fmt.Printf("%d . %s\n",i, product)

			}
			fmt.Println("Total items: ", len(cart))
		}else if choice==3{
			fmt.Println("Goodbye")
			break
		}
	}
}
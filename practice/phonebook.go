package main

import "fmt"

func phonebook() {
	phoneBook := make(map[string]string)
	var choice int
	var name, phone string

	for {
		fmt.Println("\n === Phone Book ===")
		fmt.Println("1. add contact")
		fmt.Println("2. search contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. List All")
		fmt.Println("5. Exit")
		fmt.Println("Choice: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			fmt.Println("Name: ")
			fmt.Scanln(&name)
			fmt.Println("Phone: ")
			fmt.Scanln(&phone)
			phoneBook[name]=phone
			fmt.Println("Contact Saved")
		}else if choice == 2{
			fmt.Print("Name to search :")
			fmt.Scanln(&name)
			if phone, exists:= phoneBook[name]; exists{
				fmt.Printf("%s: %s\n",name, phone)
			}else{
				fmt.Println("contact not fount")
			}
		}else if choice ==3{
			fmt.Println("name to delete: ")
			fmt.Scanln(&name)
			delete(phoneBook, name)
			fmt.Println("Deleted!")
		}else if choice==4{
			fmt.Println("\n All Contacts: ")
			for n, p:= range phoneBook{
				fmt.Printf("%s: %s\n",n, p)
			}
		}else if choice ==5{
			break
		}
	}
}
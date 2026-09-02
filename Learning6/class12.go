package main 
import "fmt"
func class12(){
	var n int 
	fmt.Println("Enter the number :")
	fmt.Scanln(&n)

	factorial:=1
	for i:=1; i<=n; i++{
		factorial=factorial*i
	}
	fmt.Printf("Factorial of %d is %d \n",n, factorial)
}
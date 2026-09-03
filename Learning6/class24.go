package main 

import "fmt"
func class24(){
	x:=10

	ptr:=&x
	value:=*ptr

	fmt.Println("x: ",x)
	fmt.Println("ptr:",ptr)
	fmt.Println("*ptr",value)

	*ptr=20
	fmt.Println("x after *ptr=20 ",x)
}
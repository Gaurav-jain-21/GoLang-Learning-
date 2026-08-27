package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	result := adder(3,5);
	fmt.Println("Result is: ",result)
	greeter();
	proRes := proAdder(2,5,8,7)
	fmt.Println("Pro result is: ",proRes)
}

func proAdder(values ...int)int{
	total:= 0
	for _,val:= range values{
		total += val
	}
	return total
}


func adder(a int,b int) int{
	return a+b;
}
func greeter(){
	fmt.Println("Namstey from golang")
}
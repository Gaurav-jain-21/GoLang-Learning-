package main

import "fmt"

// for -> only construct in go for looping
func main() {
	var  i int  =1
	for i<=5{
		var j int =1
		for j<=i{
			fmt.Print("*")
			j=j+1
		}
		fmt.Println()
		i=i+1
	}

	fmt.Println("classic for loop")

	for i:=0; i<3 ; i++ {
		fmt.Println(i)
	}
}
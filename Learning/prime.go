package main

import "fmt"

func main(){

	fmt.Println("find the prime number ")
	x:=8

	var flag bool= false

	for i:=2; i<x;i++{
		if(x%i==0){
			flag=true
			break;
		}
	}
	if(flag){
		fmt.Println("the number is not prime")
	}else{
		fmt.Println("the number is prime")
	}
}
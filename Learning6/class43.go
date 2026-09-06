package main

import "fmt"

func printAnything(v interface{}) {
	fmt.Printf("Value: %v, type : %T\n", v, v)
}
func class43(){
	printAnything(42)
	printAnything("Helllo")
}
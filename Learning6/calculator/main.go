package main

import (
	"calculator/mathutils"
	"fmt"
)

func main() {
	sum := mathutils.Add(10,5)
	diff:= mathutils.Subtract(10,5)
	prod:= mathutils.Multiply(10,5)
	quot,err:= mathutils.Divide(10,5)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot, "Error:", err)
}
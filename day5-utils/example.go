package main

import (
	"day5-utils/mathutils"
	"fmt"
)

func exmple() {
	sum := mathutils.Add(5,6)
	diff:= mathutils.Subtract(7,2)
	prod:= mathutils.Multiply(4,3)
	quot, err:= mathutils.Divide(10,5)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot, "Error:", err)
}
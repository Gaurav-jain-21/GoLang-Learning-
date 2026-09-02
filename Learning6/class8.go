package main

import "fmt"

func class8(){
	var weight float64
	var height float64
	fmt.Println("=== BMI calculator===")
	fmt.Println("Enter the you weight (kg): ")
	fmt.Scanln(&weight)
	fmt.Println("Enter you height(m): ")
	fmt.Scanln(&height)
	bmi:= weight/(height*height)

	fmt.Println("Your BMI is ", bmi)
	if bmi<18.5{
		fmt.Println("Category: underweight")
	} else if bmi < 25 {
        fmt.Println("Category: Normal weight")
    } else if bmi < 30 {
        fmt.Println("Category: Overweight")
    } else {
        fmt.Println("Category: Obese")
    }
}
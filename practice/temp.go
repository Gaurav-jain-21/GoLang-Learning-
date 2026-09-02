package main 
import "fmt"
func temp(){
	var celsius float64
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scanln(&celsius)
	fah:=(celsius*9/5)+32
	kelvin:= celsius+273.15
	fmt.Println("fahernheit: ",fah)
	fmt.Println("Kelvin ",kelvin)

}
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
func toFah(c Celsius) Fahrenheit{
	return Fahrenheit((c*9/5)+32)
}

func main(){
	var temp Celsius= 100
	f:=toFah(temp)
	fmt.Println("changing things",f)
}
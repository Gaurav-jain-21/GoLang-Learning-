package main

import "fmt"
func rectangleProperties(length, width float64)(area float64, perimeter float64, diagonal float64){
	area = length * width
	perimeter =2*(length+width)
	diagonal=(length*length +width*width)
	return 
}
func class15(){
	a,p,d:=rectangleProperties(5,3)
	fmt.Println("Area of and permiter and diag", a,p,d);
}
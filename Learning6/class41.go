package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle1 struct {
	Width, Height float64
}

func (r Rectangle1) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle1) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return  math.Pi * c.Radius*c.Radius
}

func (c Circle) Perimeter() float64{
	return 2*math.Pi*c.Radius
}
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
func class41() {
	rect := Rectangle1{Width: 10, Height: 5}
	circle:= Circle{7}

	fmt.Println(rect.Area())
	printShapeInfo(rect)
	printShapeInfo(circle)
}
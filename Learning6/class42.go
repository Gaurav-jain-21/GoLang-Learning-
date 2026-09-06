package main

import "fmt"

type Counter interface {
	Increment()
	GetValue() int
}

type SimpleCounter struct {
	value int
}

func (s *SimpleCounter) Increment() {
	s.value++
}

func (s SimpleCounter) GetValue() int {
	return s.value
}

func class42() {
	var c Counter = &SimpleCounter{}
	c.Increment()
	c.Increment()
	fmt.Println("Value", c.GetValue())
}
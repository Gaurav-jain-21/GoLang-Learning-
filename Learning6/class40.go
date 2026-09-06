package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog1 struct {
	Name string
}

func (d Dog1) Speak() string {
	return "Woof! my name is " + d.Name
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow I am" + c.Name
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello I am " + h.Name
}

func makeItTalk(s Speaker) {
	fmt.Println(s.Speak())
}

func class40(){
	dog:=Dog1{Name:"Bruno"}
	cat:=Cat{Name:"Whiskers"}
	human:= Human{Name: "Gaurav"}
    makeItTalk(dog)    // Woof! My name is Bruno
    makeItTalk(cat)    // Meow! I am Whiskers
    makeItTalk(human)
}
package main 
import "fmt"
type Animal struct{
	Name string 
}
func (a Animal) Speak(){
	fmt.Println("Some generic animal sound")

}
func (a Animal) Indro(){
	fmt.Printf("hi I am %s\n",a.Name)
}

type Dog struct{
	Animal
	Breed string
}

func (d Dog) Speak(){
	fmt.Println("woof woof")
}
func class30 (){
	dog:=Dog{
		Animal: Animal{Name: "Bruno"},
		Breed: "Golden Retriever",
	}
	dog.Indro()
	dog.Speak()
}
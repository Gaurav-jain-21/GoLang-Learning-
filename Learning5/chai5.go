package main  
import "fmt"
func chai5(){
	fmt.Println("Welcome to a class on pointers")
	var ptr  *int 
	fmt.Println("value of the empty ptr",ptr);
	myNumber :=23
	var ptr1 = &myNumber 
	fmt.Println("the pointer value of ptr1",ptr1)
	fmt.Println("value of actual pointer is ",*ptr1)

}
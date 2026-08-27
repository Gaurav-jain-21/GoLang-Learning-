package main

import "fmt"
func main(){
	fmt.Println("Variable")
	var username string = "Gaurav"
	fmt.Println(username,"is my name")
	fmt.Printf("Variable is of type : %T \n", username);

	var isLoggedIn bool= false
	fmt.Println(isLoggedIn);
	fmt.Printf("isLogged in type %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type %T \n",smallVal);
}
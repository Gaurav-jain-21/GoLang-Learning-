package main

import "fmt"
func chai2(){
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

	var smallFloat float32 = 234.534256432325
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n",smallFloat)

	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type %T \n",anotherVariable)
	
	var website  = "google.com"
	fmt.Println(website)

	numberOfUser := 300000
	fmt.Println(numberOfUser)
}
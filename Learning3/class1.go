package main

import "fmt"

const LoginToken string = "ghajkajfd" // public 

func main() {
	var username string ="gaurav"
	fmt.Println("my name is",username)
	fmt.Printf("variable is of type: %T \n",username)

	var isLoggedIn bool= true
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of isLoggedIn is %T \n",isLoggedIn)

	var smallVal int = 25
	fmt.Println(smallVal)
	fmt.Printf("the type of the smallVal is %T \n",smallVal)
	
	var smallFloat float64 = 255.9403258434
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n ", smallFloat)

	var anothervariable string
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T \n", anothervariable)

	jwtToken:=30000
	fmt.Println(jwtToken)
	fmt.Printf("the variable JwtToken %T \n ",jwtToken)

	fmt.Println(LoginToken)
	fmt.Printf("%T", LoginToken)
}
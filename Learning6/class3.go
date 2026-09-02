package main 
import "fmt"
func login(username string, password string)(bool, string){
	if username=="admin" && password=="secret123"{
		return true , "login successfull"
	}
	return false, "invalid credentials"
}

func class3(){
	success, message:= login("admin","wrongpassword")
	fmt.Printf("login status %t \n",success)
	fmt.Printf("server response: %s ", message)
}
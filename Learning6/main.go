package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Server running on http://localhost:8080")
	err:= http.ListenAndServe(":8080",nil)
	if err != nil{
		fmt.Println("Server error",err)
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Welcome to go rest api")
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if(r.Method != http.MethodGet){
		http.Error(w, "Method not allowed ",http.StatusMethodNotAllowed)
		return 
	}
	fmt.Fprintln(w,"hello from go");
}
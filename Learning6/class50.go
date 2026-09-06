package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Hello world")
}

func main(){
	http.HandleFunc("/",helloHandler)

	fmt.Println("server starting on http://localhost:8080")
	err:= http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
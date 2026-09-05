package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divided by zero")
	}
	return a/b,nil
}
func class37(){
	result,err:=divide(10,2)
	if err!=nil{
		fmt.Println("Error",err)

	}else{
		fmt.Println("Result:",result)
	}

	result2, err2:= divide(10,0)
	if err2!=nil{
		fmt.Println("Erroor",err2)
	}else{
		fmt.Println("Result",result2)
	}
}
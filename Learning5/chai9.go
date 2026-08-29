package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func chai9(){
	fmt.Println("if else in golang")

	loginCount := 23
	var result string
	if(loginCount<10){
		result= "Regular user"
	}else if(loginCount>10){
		result="Watch out"
	}else{
		result="Exactly 10 login count"
	}
	fmt.Println(result)
	fmt.Print("Enter the number to find that is even or odd : ")
	reader := bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	num , _:= strconv.ParseInt(strings.TrimSpace(input),10,64)
	if(num %2==0){
		fmt.Println("the number is even")
	}else{
		fmt.Println("the number is odd")
	}

}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func chai4(){
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("take input from the user")
	fmt.Println("Please rate our pizza between 1 and 5")
	input,_:=reader.ReadString('\n')
	number,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ", number+1)
	}

}
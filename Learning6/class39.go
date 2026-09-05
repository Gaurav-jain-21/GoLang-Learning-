package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read %s : %w", filename, err)
	}
	return string(data),nil

}

func main(){
	content, err:= readFile("missing.txt")
	if err!= nil{
		fmt.Println("Error",err)
		return 
	}
	fmt.Println(content)
}
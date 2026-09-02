package main 
import "fmt"
func class14(){
	for i:=1; i<=100; i++{
		if i%7==0 && i%13==0{
			fmt.Println("First number divisible by 7 and 13 is ",i)
			break
		}
	}
}
package main

import "fmt"

func badSwap(a,b int){
	temp:=a
	a=b
	b=temp
}

func goodSwap(a,b *int){
	temp:=*a
	*a=*b
	*b=temp
}
func class25(){
	x,y:=10,20
	badSwap(x,y)
	fmt.Println(x,y)
	goodSwap(&x,&y)
	fmt.Println(x,y)
}
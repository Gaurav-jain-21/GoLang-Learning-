package main
import "fmt"
func class17(){
	marks:=[5]int{53,63,74,77,85}
	fmt.Println("Using traditional loops: ")
	for i:=0; i<len(marks);i++{
		fmt.Println(marks[i])
	}
	fmt.Println("\n using range: ")
	for index, val:=range marks{
		fmt.Printf("marks[%d]= %d\n",index, val)
	}
}
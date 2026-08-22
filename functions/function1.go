package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages()(string, string , string){
	return "golang","java","c"
}
func main() {
	fmt.Println(add(2, 5))
	lang1,lang2,lang3:=getLanguages()
	fmt.Println(lang1,lang2,lang3)
}
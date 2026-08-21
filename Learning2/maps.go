package main

import "fmt"

func main() {
	//creating map
	m := make(map[string]string)

	//setting an element
	m["area"]="backend"
	m["name"] = "golang"
	fmt.Println(m["name"])

	// if key does not exists in the map then it returns zero value
	fmt.Println(m["phone"])

	n:=make(map[string]int)
	n["age"]=20
	fmt.Println(n["phone"])

	fmt.Println(len(n))

	delete(m,"area")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	a := map[string]int{"price":40, "phones":3}
	fmt.Println(a)

	k,ok :=m["price"]
	fmt.Println(k)
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

}
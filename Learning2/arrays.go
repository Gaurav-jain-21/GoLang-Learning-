package main

import "fmt"

func main() {
	var nums [5] int 
	// array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	var vals [4]bool
	vals[2]=true
	fmt.Println(vals)
	var name [3]string
	name[0] ="golang"
	fmt.Println(name)

	num := [3]int{1,2,4}
	fmt.Println(num)

	arr := [2][2]int{{1,2},{5,7}}
	fmt.Println(arr)
}
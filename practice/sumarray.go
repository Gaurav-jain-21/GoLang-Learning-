package main

import "fmt"

func sunSlice(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum = sum + val
	}
	return sum
}
func sumarray() {
	numbers := []int{1, 2, 3, 5, 5}
	result := sunSlice(numbers)
	fmt.Println(result)
}
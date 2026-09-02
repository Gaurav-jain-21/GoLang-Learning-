package main

import "fmt"

func reverse(numbers []int) []int {
	start := 0
	end := len(numbers)-1
	for start < end {
		temp := numbers[start]
		numbers[start] = numbers[end]
		numbers[end] = temp

		start = start + 1
		end = end - 1
	}
	return numbers
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(numbers))
}
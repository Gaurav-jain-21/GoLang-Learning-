package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(len(nums))
	var num = make([]int, 2)
	fmt.Println(num)

	num= append(num,1)
	fmt.Println(num)
	num=append(num,2)
	num=append(num,4)
	num=append(num,7)
	fmt.Println(num)
	fmt.Println(cap(num))
	fmt.Println(len(num))
	var arr =make([]int,0,5)
	fmt.Println(len(arr),"length of the slice")
	fmt.Println(cap(arr), "capacity of the slice")
	for i:=0; i<5;i++{
		arr= append(arr,i)
	}
	fmt.Println(arr)

	var num1 =[]int{8,5,2,1}
	fmt.Println(num1[:2])

	var nums2 = []int{1,2}
	var nums3 = []int{1,2}
	fmt.Println(slices.Equal(nums2, nums3))
	
}
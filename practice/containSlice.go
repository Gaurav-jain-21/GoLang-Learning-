package main

import "fmt"

func contains(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
func containSlice() {
	items := []string{"laptop", "tv", "phone", "tablet"}
	find := contains(items, "tv")
	fmt.Println("is tv in the list ", find)
}
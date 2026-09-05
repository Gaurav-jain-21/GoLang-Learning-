package main

import (
	"fmt"
	"os"
)

func class33() {
	content := []byte("Hello world this is a new file")
	err := os.WriteFile("data.txt", content, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File written successfully")
}
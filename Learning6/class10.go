package main

import "fmt"

func class10(){
	var marks int
    fmt.Print("Enter your marks (0-100): ")
    fmt.Scanln(&marks)

    if marks >= 90 {
        fmt.Println("Grade: A+ 🌟")
        fmt.Println("Outstanding performance!")
    } else if marks >= 80 {
        fmt.Println("Grade: A")
        fmt.Println("Excellent work!")
    } else if marks >= 70 {
        fmt.Println("Grade: B")
        fmt.Println("Good job, keep improving!")
    } else if marks >= 60 {
        fmt.Println("Grade: C")
        fmt.Println("You passed, but study harder!")
    } else if marks >= 40 {
        fmt.Println("Grade: D")
        fmt.Println("Barely passed. Need serious work!")
    } else {
        fmt.Println("Grade: F ❌")
        fmt.Println("Failed. Please reappear for the exam.")
    }
}
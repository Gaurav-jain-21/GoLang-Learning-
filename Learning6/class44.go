package main

import "fmt"

func describe(i interface{}) {
    // Type assertion: try to extract as string
    str, ok := i.(string)
    if ok {
        fmt.Println("It's a string:", str)
        return
    }
    
    // Try as int
    num, ok := i.(int)
    if ok {
        fmt.Println("It's an int:", num)
        return
    }
    
    fmt.Println("Unknown type")
}

func class44() {
    describe("Gaurav")
    describe(100)
    describe(3.14)
}
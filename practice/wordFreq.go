package main

import (
	"fmt"
	"strings"
)

func wordFreq(){
	sentence:= "go id un and go is powerfull"
	words:= strings.Fields(sentence)
	wordCount:= make(map[string]int)

	for _, word:=range words{
		wordCount[word]++
	}
	for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
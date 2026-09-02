package main

import "fmt"
func class16(){
	var scores [5]int
	scores[0]=85
	scores[1]=90
	scores[2]=83
	scores[3]=75
	scores[4]=91
	fmt.Println("Scores:",scores)
	fmt.Println("First scores : ",scores[0])
	fmt.Println("last Scores : ",scores[4])
	fmt.Println("Number of subjects: ",len(scores))

	temperatures:=[4]float64{3.64,4.642,74.4,234.54}
	fmt.Println("week temps: ",temperatures)

	names:=[...]string{"Gaurav","Rahul","Priya"}
	fmt.Println("Names: ",names)
	fmt.Println("count ",len(names))
}
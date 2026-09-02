package main

import "fmt"
func calcuteGrade(percentage float64) string{
	if percentage>=90{
		return "A+"
	}else if percentage>=80{
		return "A"
	}else if percentage>=70{
		return "B+"
	}else if percentage>=60{
		return "B"
	}else{
		return "F"
	}
}

func grade(){
	student:= map[string][]int{
		"Gaurav":{90, 90, 99, 92, 88},
		"Rahul":{75, 80, 85, 70, 90},
	}

	for name, marks:=range student{
		total:=0
		for _ , mark:= range marks{
			total+=mark
		}
		percentage:= float64(total)/float64(len(marks))
		fmt.Println("percentage is ",percentage)
		grade:=calcuteGrade(percentage)
		fmt.Println("name of the with is ",name, grade)
	}
}
package main

import "fmt"

type Person struct{
	Name string 
	Age int
	Address string 
}

type Student struct{
	Person
	RollNumber string 
	Class string 
	Marks []int
}

func (s Student) GetAverage() float64{
	if len(s.Marks)==0{
		return 0
	}
	total:=0
	for _,mark:=range s.Marks{
		total+=mark
	}
	return float64(total)/float64(len(s.Marks))
}
func (s Student) DisplayReport(){
	avg:= s.GetAverage()
	fmt.Printf("\n====Report Card===\n")
	fmt.Printf("Name: %s\n",s.Name)
	fmt.Printf("Age: %d\n",s.Age)
	fmt.Printf("Roll: %s\n",s.RollNumber)
	fmt.Printf("Class: %s\n",s.Class)
	fmt.Printf("Marks: %v\n",s.Marks)
	fmt.Printf("Average: %.2f\n",avg)

	if avg >= 40 {
        fmt.Println("Status: PASSED ✅")
    } else {
        fmt.Println("Status: FAILED ❌")
    }
}

func studentmanga(){
	students:=[]Student{
		{
            Person:     Person{Name: "Gaurav", Age: 22, Address: "Mumbai"},
            RollNumber: "R001",
            Class:      "12th",
            Marks:      []int{85, 90, 78, 92, 88},
        },
        {
            Person:     Person{Name: "Rahul", Age: 21, Address: "Delhi"},
            RollNumber: "R002",
            Class:      "12th",
            Marks:      []int{35, 40, 38, 42, 36},
        },
	}
	for _,student:=range students{
		student.DisplayReport()
	}
}
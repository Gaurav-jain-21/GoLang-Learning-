package main 

import "fmt"

type Applicants struct{
	Id int 
	Name string 
	YearOfExperience int
}

func evaluateApplicat(app Applicants)(bool,string){
	if app.YearOfExperience<2{
		return false, "Need more Experience"
	}
	return true, "Shortlisted for Interview"
}

func class5(){
	applicant:=Applicants{
		Id: 101,
		Name: "Gaurav",
		YearOfExperience: 3,
	}
	approved,message:= evaluateApplicat(applicant)
	fmt.Printf("the applicant is %t and message is %s ",approved,message)

}
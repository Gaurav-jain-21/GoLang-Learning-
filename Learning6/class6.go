package main

import "fmt"

type Applicant struct {
	ID                int
	Name              string
	YearsOfExperience int
	Status            string
}

func (a *Applicant) Hire() {
	a.Status = "Hired"
}
func class6() {
	applicant := Applicant{
		ID:               101,
		Name:             "Gaurav",
		YearsOfExperience: 3,
		Status:           "Pending",
	}
	fmt.Printf("Status before interview: %s\n", applicant.Status)
	applicant.Hire()
	fmt.Printf("Status after interview: %s\n", applicant.Status)
}
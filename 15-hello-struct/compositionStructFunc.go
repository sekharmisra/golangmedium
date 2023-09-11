package main

import "fmt"

type Details struct {
	name       string
	age        int
	address    string
	postalcode string
}

type Student struct {
	college      string
	branch       string
	emailaddress string
	rank         int
	Details
}

func compositionStructFunc() {
	details := Details{
		name:       "Sekhar Misra",
		age:        40,
		address:    "Waterloo",
		postalcode: "7",
	}

	studentDetails := Student{
		college:      "NIT Durgapur",
		branch:       "Electronics and Communication Engineering",
		emailaddress: "test@test.com",
		rank:         30,
		Details:      details,
	}

	studentDetails.StudentDetails()
}

func (details Details) showDetails() {
	fmt.Println("Name", details.name)
	fmt.Println("Age", details.age)
	fmt.Println("Address", details.address)
	fmt.Println("Postalcode", details.postalcode)
}

func (student Student) StudentDetails() {
	fmt.Println("College", student.college)
	fmt.Println("Branch", student.branch)
	fmt.Println("Email Address", student.emailaddress)
	fmt.Println("Rank", student.rank)
	student.showDetails()
}

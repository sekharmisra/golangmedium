package main

import "fmt"

func promotedStructFunc() {
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

	studentDetails := Student{
		college:      "NIT Durgapur",
		branch:       "Electronics and Communication Engineering",
		emailaddress: "test@test.com",
		rank:         30,
		Details: Details{
			name:       "Sekhar Misra",
			age:        40,
			address:    "Waterloo",
			postalcode: "7",
		},
	}

	fmt.Println(studentDetails)
}

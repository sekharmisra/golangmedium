package main

import "fmt"

func simpleStructFunc() {
	type Details struct {
		name       string
		age        int
		address    string
		postalcode string
	}

	user_details := Details{
		name:       "Sekhar Misra",
		age:        40,
		address:    "Waterloo",
		postalcode: "7",
	}

	fmt.Println(user_details)
}

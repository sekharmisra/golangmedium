package main

import "fmt"

func switchGo() {

	var schedule string = "03Sep"

	switch schedule {
	case "01Sep":
		fmt.Println("01Sep - Kafka")
	case "02Sep":
		fmt.Println("Git and GitHub")
	case "03Sep":
		fmt.Println("Go Lang Basics!")
	default:
		fmt.Println("Doing nothing!")
	}
}

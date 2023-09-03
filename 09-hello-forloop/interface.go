package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
	skills string
}

func interfaceExample() {
	fmt.Println("******Interface in Go STart*************")
	emp := employee{
		name:   "Sekhar Misra",
		age:    40,
		salary: 100,
		skills: "GoLang!",
	}
	printType(emp)
	fmt.Println("******Interface in Go End*************")
}

func printType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Decimal")
	case employee:
		fmt.Println("Employee")
	default:
		fmt.Println("Unknown type")
	}
}

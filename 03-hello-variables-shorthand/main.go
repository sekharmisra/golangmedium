package main

import "fmt"

func main() {

	//Shorthand delcaration in Go!
	name := "Sekhar Misra"
	age := 40
	weight := 100.6
	isMotivated := true

	fmt.Printf("My name is %v & type is %T\n", name, name)
	fmt.Printf("My age is %v & type is %T\n", age, age)
	fmt.Printf("My weight is %v & type is %T\n", weight, weight)
	fmt.Printf("Am I motivated? %v & type is %T\n", isMotivated, isMotivated)

}

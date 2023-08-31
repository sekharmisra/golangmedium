package main

import "fmt"

func main() {

	//Explicit declarations of variables
	var name string = "Sekhar Misra"
	var age int = 40
	var weight float64 = 100.6
	var isMotivated bool = true

	fmt.Printf("My name is %v & type is %T\n", name, name)
	fmt.Printf("My age is %v & type is %T\n", age, age)
	fmt.Printf("My weight is %v & type is %T\n", weight, weight)
	fmt.Printf("Am I motivated? %v & type is %T\n", isMotivated, isMotivated)

	//one more example for completeness, this is below instead of shorthand delcaration but still implcitly typed by assiging value!
	var helloString = "Hello String!"
	fmt.Printf("Hello string value is: %v & type is %T\n", helloString, helloString)

}

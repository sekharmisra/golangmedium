package main

import "fmt"

func main() {
	numx := 10
	numy := 11

	fmt.Println(addNumFunc(numx, numy))
	fmt.Println(addStringFunc("Sekhar", "Misra"))
	//fmt.Println(addVariadicFunc(addSlice()))

	nameSlice := []string{}
	nameSlice = addSlice()
	fmt.Println(addVariadicFunc(nameSlice))
	fmt.Println(averageSalary(3125, 2018, 2511, 3125))

	//anonymous function
	greet := func() {
		fmt.Println("Hello World from Anonymous Function")
	}

	greet()
	fmt.Printf("Type of greet is : %T", greet)

}

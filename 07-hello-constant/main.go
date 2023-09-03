package main

import "fmt"

func main() {

	//function level constants
	const (
		wish = "hello world!" //implicity typed constant
		age  = 10.10          //implicity typed constant
	)
	const isMotivated bool = true

	if isMotivated {
		const motivated string = "motivated!"
		fmt.Printf("Motivated %v %T\n", motivated, motivated)
	}

	fmt.Printf("Wish %v %T\n", wish, wish)
	fmt.Printf("Age %v %T\n", age, age)
	fmt.Printf("IsMotivated %v %T\n", isMotivated, isMotivated)
	fmt.Printf("Number %v %T\n", num, num)

}

const num int = 100 // Package level constant

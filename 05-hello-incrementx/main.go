package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("The value of x is %v\n", x)

	x = +10
	fmt.Printf("The value of x is %v\n", x)

	x += 10
	fmt.Printf("The value of x is %v\n", x)

	x++
	fmt.Printf("The value of x is %v\n", x)

	x--
	fmt.Printf("The value of x is %v\n", x)

}

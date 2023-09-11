package main

import "fmt"

func nestedForLoop() {

	fmt.Println("*****************Nested For Loop Start*****************")
	for i := 0; i < 10; i++ {
		fmt.Printf("Value of i: %v\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("Value of j: %v\n", j)
		}
	}
	fmt.Println("*****************Nested For Loop End*****************")
}

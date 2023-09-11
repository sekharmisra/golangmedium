package main

import "fmt"

func oddEvenCompute() {
	fmt.Println("******Odd Even Go Start*************")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("\nValue %v is Even\n", i)
		} else {
			fmt.Printf("\nValue %v is Odd\n", i)
		}
	}
	fmt.Println("******Odd Even Go Start*************")
}

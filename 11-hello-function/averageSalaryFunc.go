package main

import "fmt"

//So called variadic function(s)
func averageSalary(elements ...float64) float64 {
	fmt.Printf("\nElements into the slice are as follows:\n %v\n", elements)
	fmt.Printf("\nType %T\n", elements)

	//Loop through each element and add the value to temp variable
	var sum float64
	for _, val := range elements {
		sum += val
	}
	return sum / float64(len(elements))
}

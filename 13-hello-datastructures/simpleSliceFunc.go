package main

import "fmt"

func simpleSliceFunc() {
	//Slice declaration
	//var simpleSlice []int
	simpleSlice := []int{}

	for i := 0; i < 5; i++ {
		simpleSlice = append(simpleSlice, i)
	}
	simpleSlice = append(simpleSlice, 6)

	for index := 0; index < len(simpleSlice); index++ {
		fmt.Printf("\nValue at Index %v is %v\n", index, simpleSlice[index])
	}
}

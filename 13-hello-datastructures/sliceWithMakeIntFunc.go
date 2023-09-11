package main

import "fmt"

func sliceWithMakeIntFunc() {
	var my_slice = make([]int, 2, 5)

	my_slice[0] = 1
	my_slice[1] = 2
	fmt.Printf("\nSize of slice is %v and capacity is %v\n", len(my_slice), cap(my_slice))

	for i := 3; i < 15; i++ {
		my_slice = append(my_slice, i)
		fmt.Printf("\nSize of slice is %v and capacity is %v\n", len(my_slice), cap(my_slice))
	}
}

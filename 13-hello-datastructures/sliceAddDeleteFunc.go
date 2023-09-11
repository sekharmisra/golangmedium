package main

import "fmt"

func sliceAddDeleteFunc() {
	var my_slice1 = []string{"eat", "sleep"}
	var my_slice2 = []string{"gym", "work"}

	//appending one slice to another
	my_slice1 = append(my_slice1, my_slice2...)
	fmt.Println(my_slice1)

	//deleting elements from the slice!
	fmt.Println(my_slice1[0:1])
	fmt.Println(my_slice1[2:])
	my_slice1 = append(my_slice1[0:1], my_slice1[2:]...)
	fmt.Println(my_slice1)
}

package main

import "fmt"

func sliceWithMakeFunc() {

	//Slice with make function

	var slice = make([]string, 2, 5)

	slice = append(slice, "Tea")
	slice = append(slice, "Cofee!")
	slice = append(slice, "Breakfast")

	fmt.Println(slice)
}

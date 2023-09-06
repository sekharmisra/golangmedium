package main

import "fmt"

func simpleStringSlice() {
	//String slice
	var simpleStringSlice = []string{"Tea", "Cofee", "Breakfast"}
	fmt.Println("Slice data", simpleStringSlice)

	for _, val := range simpleStringSlice {
		fmt.Println(val)
	}
}

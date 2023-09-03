package main

import "fmt"

func main() {
	numx := 10
	numy := 11

	fmt.Println(addNumFunc(numx, numy))
	fmt.Println(addStringFunc("Sekhar", "Misra"))
	//fmt.Println(addVariadicFunc(addSlice()))

	nameSlice := []string{}
	nameSlice = addSlice()
	fmt.Println(addVariadicFunc(nameSlice))

}

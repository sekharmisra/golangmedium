package main

import "fmt"

func main() {

	//Example of Array - arrayFunc, Just increase the element from 5 to 6 see index out of range error!
	fmt.Println("\nArray function not much used in Go!\n")
	arrayFunc()

	//Simple slice function
	fmt.Println("\nSimple slice example!\n")
	simpleSliceFunc()

	//Simple string slice function
	fmt.Println("\nString Simple slice example!\n")
	simpleStringSlice()

	//Call arry slice reference function
	fmt.Println("\nSlice with array reference function example!\n")
	arraySliceRefFunc()

	//Slice with make function!
	fmt.Println("\nSlice with make function example!\n")
	sliceWithMakeFunc()

	//Calling sliceWithMakeIntFunc()
	fmt.Println("\nMore complex example!\n")
	sliceWithMakeIntFunc()

	//Calling sliceAddDeleteFunc
	fmt.Println("\nsliceAddDeleteFunc!\n")
	sliceAddDeleteFunc()

	//twodminSliceFunc
	fmt.Println("\ntwodminSliceFunc!\n")
	twodminSliceFunc()
}

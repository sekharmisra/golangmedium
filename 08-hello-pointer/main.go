package main

import "fmt"

func main() {

	var name string = "Sekhar Misra"
	var ptrName *string = &name //ptrName points to memory address of name

	//Print vlaues
	fmt.Printf("Value of name is %v\n", name)
	fmt.Printf("Value of memory of name is %v\n", ptrName)
	fmt.Printf("Value of pointer is %v\n", *ptrName)

	//Let us change the name
	*ptrName = "Kunjan Misra"

	//Print the values again
	fmt.Printf("Value of name is %v\n", name)
	fmt.Printf("Value of memory of name is %v\n", ptrName)
	fmt.Printf("Value of pointer is %v\n", *ptrName)
	fmt.Printf("Memory of ptrName is %v\n", &ptrName)

	//Let us change the name again!
	*ptrName = "Aarna Misra"

	//Print the values again
	fmt.Printf("Value of name is %v\n", name)
	fmt.Printf("Value of memory of name is %v\n", ptrName)
	fmt.Printf("Value of pointer is %v\n", *ptrName)
	fmt.Printf("Memory of ptrName is %v\n", &ptrName)
}

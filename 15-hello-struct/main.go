package main

import "fmt"

func main() {
	//Calling simpleStructFunc
	fmt.Println("*****Calling simpleStructFunc***********")
	simpleStructFunc()

	//Calling promotedStructFunc
	fmt.Println("*****Calling promotedStructFunc***********")
	promotedStructFunc()

	//Calling compositionStructFunc
	fmt.Println("*****Calling compositionStructFunc***********")
	compositionStructFunc()
}

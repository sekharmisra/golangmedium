// packages are collection of files arranged in single package. You need to init go module before compiling the program!
package main

//importing single package to use println function!
import "fmt"

//main() is the main entry point of the program!
func main() {
	fmt.Println("Hello World! Today is Aug 31 2023. I am learning Go from sctrach!")
	myVariable := "Hello World from variable!"

	fmt.Printf("Here I am printing value from the implicitly typed variable! ***** %v\n", myVariable)
	fmt.Printf("Here I am printing value of the type of myVariable! ***** %T\n", myVariable)

	//Calling func
	hellouuid()

}

package main

import "fmt"

func main() {
	defer sayHi()
	sayBye()
}

func sayHi() {
	fmt.Println("Say Hi!!")
}

func sayBye() {
	fmt.Println("Say Bye!!")
}

package main

import "fmt"

func simpleMapFunc() {
	my_map := map[int]string{
		1: "Animal",
		2: "Human being",
		3: "Alien",
	}
	fmt.Println(my_map)

	//check for existing key
	val, ok := my_map[3]
	if !ok {
		fmt.Println("Key not present")
	} else {
		fmt.Println(val)
		fmt.Println("Key going to be deleted!")
		delete(my_map, 3)
	}
	//check for existing key
	val1, ok1 := my_map[3]
	if !ok1 {
		fmt.Println("Key not present")
	} else {
		fmt.Println(val1)
	}

}

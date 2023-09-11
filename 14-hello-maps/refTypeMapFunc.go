package main

import "fmt"

func refTypeMapFunc() {
	//Demonstrating Map is reference type.Copying one map into another still references same data structures and changes are relected into each other.
	my_map := make(map[int]string)
	my_map[0] = "Animal"
	my_map[1] = "Human being"
	my_map[2] = "Alien"

	fmt.Println("Original Map: ", my_map)
	my_new_map := my_map

	my_new_map[3] = "Moon"
	fmt.Println("Original Map: ", my_map)
	fmt.Println("New Map: ", my_new_map)
}

package main

import "fmt"

func makeMapFunc() {
	my_map := make(map[int]string)
	my_map[0] = "Animal"
	my_map[1] = "Human being"
	my_map[2] = "Alien"

	for key, val := range my_map {
		fmt.Printf("\n%v:%v\n", key, val)
	}
}

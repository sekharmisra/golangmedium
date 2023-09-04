package main

import (
	"fmt"
)

func printNumbers(data []int, callbackFunc func(int)) {
	for _, val := range data {
		callbackFunc(val)
	}
}

func callbackMainFunc() {
	data := []int{1, 2, 3, 4, 5}
	printNumbers(data, func(val int) {
		fmt.Println(val)
	})
}

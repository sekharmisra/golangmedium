package main

import "fmt"

func arrayFunc() {
	var arr [5]int

	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	for index := 0; index < len(arr); index++ {
		fmt.Printf("\nValue at Index %v is %v\n", index, arr[index])
	}
}

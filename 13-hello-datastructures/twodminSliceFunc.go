package main

import "fmt"

func twodminSliceFunc() {

	//Two dimensional slice
	var my_slice = make([][]string, 0)

	user1 := make([]string, 3)
	user1[0] = "Dummy user 1"
	user1[1] = "40"
	user1[2] = "Coding!"

	my_slice = append(my_slice, user1)

	user2 := make([]string, 3)
	user2[0] = "Dummy user 2"
	user2[1] = "40"
	user2[2] = "Coding!"

	my_slice = append(my_slice, user2)

	fmt.Println(my_slice)

}

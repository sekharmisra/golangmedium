package main

import "fmt"

func arraySliceRefFunc() {
	//Declare array of size 5 with string and elements into it
	arr := [5]string{"Tea", "Cofee", "Breakfast", "School", "Gym"}

	//Assign array to the slice
	my_slice := arr[1:]

	//Print slice values
	for _, val := range my_slice {
		fmt.Println(val)
	}

	//update array elements value from Gym to Workout

	arr[4] = "Workout"

	//See if the slice changes to prove that slice is a reference to an array!
	for _, val := range my_slice {
		fmt.Println(val)
	}
}

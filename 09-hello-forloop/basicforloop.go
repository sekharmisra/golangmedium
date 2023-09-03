package main

import "fmt"

func firstForLoop() {
	fmt.Println("******************First For Loop Start!******************")
	for i := 0; i < 10; i++ {
		fmt.Println((i))
	}
	fmt.Println("******************First For Loop End!******************")
}

func secondForLoop() {
	fmt.Println("******************Second For Loop Start!******************")
	//Same result
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("******************Second For Loop End!******************")
}

func thirdForLoop() {
	fmt.Println("******************Thrid For Loop Start!******************")
	i := 0
	//infite for loop until break
	for {
		fmt.Println(i)
		i++

		if i == 10 {
			break
		}
	}
	fmt.Println("******************Thrid For Loop End!******************")
}

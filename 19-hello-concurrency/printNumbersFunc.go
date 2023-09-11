package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i <= 50; i++ {
		fmt.Printf("Data func1:%v\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}

func printAnotherSeries() {
	for i := 51; i <= 100; i++ {
		fmt.Printf("Data func2:%v\n", i)
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}

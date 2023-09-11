package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go printNumbers()
	go printAnotherSeries()
	wg.Wait()

	wg2.Add(2)
	go printNumbersRaceCondition("foo")
	go printNumbersRaceCondition("bar")
	wg2.Wait()
	fmt.Printf("Final value of counter is %v\n", counter)
}

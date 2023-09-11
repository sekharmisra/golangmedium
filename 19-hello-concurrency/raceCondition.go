package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg2 = sync.WaitGroup{}
var mutex sync.Mutex

func printNumbersRaceCondition(val string) {
	for i := 0; i <= 25; i++ {
		mutex.Lock() //prevents race condition, remove to see one
		x := counter
		x++
		time.Sleep(100 * time.Millisecond)
		counter = x
		fmt.Printf("%v has counter %v\n", val, counter)
		mutex.Unlock() //prevents race condition, remove to see one
	}
	wg2.Done()

	//check using go run -race source_file.go
	//-race requires cgo; enable cgo by setting CGO_ENABLED=1
}

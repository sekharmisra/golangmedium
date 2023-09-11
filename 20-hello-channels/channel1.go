package main

import (
	"fmt"
	"time"
)

func channel1() {
	ch := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(1000 * time.Millisecond)
}

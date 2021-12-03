package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) //worker number 2

	go func() {
		// worker 1 do something
		time.Sleep(10 * time.Second)
		fmt.Println("goroutine 1 done！")
		wg.Done()
	}()

	go func() {
		// worker 2 do something
		fmt.Println("goroutine 2 done！")
		wg.Done()
	}()

	wg.Wait() // wait all waiter done
	fmt.Println("all work done！")
}

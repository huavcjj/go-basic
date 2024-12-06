package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()

	go func() {
		time.Sleep(time.Hour)
	}()

	wg.Wait()

	// ch <- 1
	// time.Sleep(time.Second)
}

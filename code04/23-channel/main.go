package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 100)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	mc := make(chan struct{}, 0)
	go func() {
		sum := 0
		for {
			a, ok := <-ch
			if !ok {
				break
			}
			sum += a
		}
		fmt.Printf("sum: %d\n", sum)
		mc <- struct{}{}
	}()

	wg.Wait()
	close(ch)

	<-mc

}

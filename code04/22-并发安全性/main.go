package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	n int32
	// lock = sync.Mutex{}
)

func foo() {
	for i := 0; i < 100000; i++ {
		// lock.Lock()
		// n++
		// lock.Unlock()
		atomic.AddInt32(&n, 1)

	}
	fmt.Printf("n= %d\n", n)
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		foo()
	}()
	go func() {
		defer wg.Done()
		foo()
	}()
	wg.Wait()
	fmt.Printf("n= %d\n", n)

}

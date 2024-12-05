package main

import (
	"fmt"
	"sync"
	"time"
)

func foo() string {
	str := "Hello, World!"

	defer fmt.Println("deferred call in foo 1", str)
	fmt.Println(str)
	defer fmt.Println("deferred call in foo 2", str)
	defer func() {
		fmt.Println("deferred call in foo 3", str)
	}()
	str = "Hello, Go!"
	return str

}

func hello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello!")
}

func world() {
	fmt.Println("World!")
	go hello()
}

func f1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f1", r)
		}
	}()
	a, b := 3, 0
	fmt.Println(a, b)
	_ = a / b
	fmt.Println("f1 done")
}

func main() {

	m := map[string]int{"apple": 100, "banana": 200}

	if v, ok := m["apple"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	foo()

	start := time.Now()
	fmt.Println(start.Format("2006-01-02 15:04:05"))
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed.Microseconds())

	duration := time.Duration(10) * time.Second
	end := start.Add(duration)
	fmt.Println(end.Format("2006-01-02 15:04:05"))

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2")
	}()
	wg.Wait()

	go world()

	go f1()
	time.Sleep(3 * time.Second)
	fmt.Println("main done")

}

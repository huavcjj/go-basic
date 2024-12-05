package main

import (
	"fmt"
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

}

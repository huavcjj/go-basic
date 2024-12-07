package main

import (
	"fmt"
	"slices"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	values := []string{"a", "b", "c"}
	wg.Add(len(values))

	for _, v := range values {
		fmt.Printf("%p\n", &v)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}

	wg.Wait()

	for i := range 3 {
		fmt.Println(i)
	}

	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}

	// merged := append(append(arr1, arr2...), arr3...)
	// fmt.Println(merged)

	merged := slices.Concat(arr1, arr2, arr3)
	fmt.Println(merged)
}

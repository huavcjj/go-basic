package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)    // [1 2] 3
	sum(1, 2, 3) // [1 2 3] 6

	// 複数の引数をスライスとして持っているなら、以下の func(slice...) という形式で可変長引数関数にスライスの中身を渡せる
	nums := []int{1, 2, 3, 4}
	sum(nums...) // [1 2 3 4] 10
}

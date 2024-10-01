package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i) //1 2 3
		i = i + 1
	}

	// 初期化、条件、ループ間処理を書く古典的な for ループ
	for j := 7; j <= 9; j++ {
		fmt.Println(j) //7 8 9
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n) //1 3 5
	}
}

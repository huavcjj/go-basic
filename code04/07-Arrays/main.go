package main

import "fmt"

func main() {

	var a [5]int           //配列を宣言すると、要素はゼロ値（int 型の場合は0）になる
	fmt.Println("emp:", a) //emp: [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a)    //set: [0 0 0 0 100]
	fmt.Println("get:", a[4]) //get: 100

	fmt.Println("len:", len(a)) //len: 5

	//このように書けば、配列の宣言と初期化を一行で済ませられる
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) //dcl: [1 2 3 4 5]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) //2d:  [[0 1 2] [1 2 3]]
}

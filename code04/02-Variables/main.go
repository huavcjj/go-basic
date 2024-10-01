package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a) //initial

	var b, c int = 1, 2
	fmt.Println(b, c) //1 2

	var d = true
	fmt.Println(d) //true

	var e int
	fmt.Println(e) //0
	//! 明示的に初期化していない変数の値はゼロ値になる

	f := "apple"
	// var f string = "apple"
	fmt.Println(f) //apple
}

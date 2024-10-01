package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) //constant

	const n = 50000000 //! const 文は var 文が書けるところならどこでも書ける

	const d = 3e20 / n
	fmt.Println(d) //6e+11

	//! 数値の定数は、明示的にキャストするなどして型が決まるまでは、型を持たない。
	fmt.Println(int64(d)) //600000000000

	fmt.Println(math.Sin(n)) //-0.28470407323754404
	//! 数値の型はそれを必要とするコンテキストによって与えられる。
}

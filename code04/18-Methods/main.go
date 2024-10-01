package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// この area メソッドは、 *rectをレシーバ型として持つ
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 構造体に定義されたメソッドを呼び出す
	fmt.Println("area: ", r.area())  // area:  50
	fmt.Println("perim:", r.perim()) // perim: 30

	// Go は、メソッドを呼び出す際に、値とポインタの変換を自動的に行う
	//メソッドを呼び出す時の構造体の不要なコピーを避けたり、元の構造体自体を変更したりするためには、ポインタレシーバを使う
	rp := &r
	fmt.Println("area: ", rp.area())  // area:  50
	fmt.Println("perim:", rp.perim()) // perim: 30
}

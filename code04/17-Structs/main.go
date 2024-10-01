package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20}) // {Bob 20}

	//構造体を初期化するときに、 フィールド名を指定することもできる
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

	//省略されたフィールドはゼロ値になる
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	fmt.Println(newPerson("Jon")) // &{Jon 42}

	//ドットを使ってフィールドにアクセスする
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean

	//構造体のポインタにもドットが使える
	sp := &s
	fmt.Println(sp.age) // 50

	//構造体は変更可能 (mutable)
	sp.age = 51
	fmt.Println(sp.age) // 51
}

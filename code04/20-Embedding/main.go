package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	//リテラルを使って構造体を作成する場合には、埋め込まれた型を明示的に初期化する必要がある
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	//co に対して、co.num のように base のフィールドに直接アクセスできる
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // co={num: 1, str: some name}

	//代わりに、埋め込まれた型の名前を使って、 フルパスで指定することもできる
	fmt.Println("also num:", co.base.num) // also num: 1

	//埋め込まれた型のメソッドも呼び出すことができる
	fmt.Println("describe:", co.describe()) // describe: base with num=1

	type describer interface {
		describe() string
	}

	//メソッドをもつ構造体の埋め込みは、他のインターフェースを埋め込むこともできる
	var d describer = co
	fmt.Println("describer:", d.describe()) // describer: base with num=1
}

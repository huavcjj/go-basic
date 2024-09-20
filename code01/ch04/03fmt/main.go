package main

import (
	"fmt"
)

type F struct {
	Name string
	Age  int
}

func main() {
	x := 1
	s := fmt.Sprintf("%d", x)
	fmt.Println(s)

	// f, err := os.Create("output.dat")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// fmt.Fprintf(f, "x=%d\n", x)

	f := &F{"Alice", 20}
	fmt.Printf("%v\n", f)  // &{Alice 20}
	fmt.Printf("%+v\n", f) // &{Name:Alice Age:20}
	fmt.Printf("%#v\n", f) // &main.F{Name:"Alice", Age:20}
	fmt.Printf("%T\n", f)  // *main.F
	fmt.Printf("%T\n", *f) // main.F

}

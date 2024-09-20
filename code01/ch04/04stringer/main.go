package main

import "fmt"

// type Stringer interface {
// 	String() string
// }

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME: %s, AGE: %d", f.Name, f.Age)
}

func main() {
	f := &F{"Alice", 20}
	fmt.Printf("%v\n", f) // NAME: Alice, AGE: 20

}

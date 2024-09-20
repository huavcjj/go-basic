package main

import "fmt"

type Fruit int

const (
	Apple  Fruit = iota // 0
	Orange              // 1
	Banana              // 2
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Orange:
		return "Orange"
	case Banana:
		return "Banana"
	}
	return fmt.Sprintf("Fruit(%d)", i)
}

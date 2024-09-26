package main

import (
	"errors"
	"fmt"
)

func main() {
	err01 := errors.New("something wrong")
	err02 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	fmt.Println(err01.Error())
	fmt.Println(err01 == err02)

	err0 := fmt.Errorf("add info : %w", errors.New("original error"))
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	fmt.Println(errors.Unwrap(err0))
}

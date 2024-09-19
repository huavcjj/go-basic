package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 引数を割り算し浮動小数点の値として結果を表示するプログラム
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor must not be 0")
	}
	return a / b, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "number of arguments must be 2")
		os.Exit(1)
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "first argument must be float value: %v\n", err)
		os.Exit(1)
	}
	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "second argument must be float value: %v\n", err)
		os.Exit(1)
	}
	result, err := divide(a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid argument: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

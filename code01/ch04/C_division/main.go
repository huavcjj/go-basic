package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 引数を割り算し浮動小数点の値として結果を表示するプログラム
// 短くしたプログラム(エラーを十分に伝える)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor must not be 0")
	}
	return a / b, nil
}

func exitIf(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "number of arguments must be 2")
		os.Exit(1)
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	exitIf(err, "first argument must be float value")
	b, err := strconv.ParseFloat(os.Args[2], 64)
	exitIf(err, "second argument must be float value")
	result, err := divide(a, b)
	exitIf(err, "invalid argument")
	fmt.Println(result)
}

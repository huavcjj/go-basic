package calculator

import "fmt"

var offest float64 = 1
var Offest float64 = 1

func Sum(a, b float64) float64 {
	fmt.Println("multiply:", multiply(a, b))
	fmt.Println("Multiply:", Multiply(a, b))
	return a + b + offest
}

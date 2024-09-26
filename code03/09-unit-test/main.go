package main

func Add(x, y int) int {
	return x + y
}
func Divide(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}
func main() {
	// x, y := 3, 5
	// fmt.Printf("%v %v\n", Add(x, y), Divide(x, y))

	//go test -v -cover -coverprofile=coverage.out .
	//go tool cover -html=coverage.out

}

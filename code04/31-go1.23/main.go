package main

import (
	"fmt"
	"maps"
	"slices"
)

func slice() {
	arr := []int{1, 2, 3}
	brr := slices.Repeat(arr, 3)
	fmt.Println(brr)

	slices.Sort(brr)
	fmt.Println(brr)

	slices.Reverse(brr)
	fmt.Println(brr)

}

func sortMap() {
	m := map[string]struct{}{
		"foo": {},
		"bar": {},
		"baz": {},
	}
	for _, k := range slices.Sorted(maps.Keys(m)) {
		fmt.Printf("%s\n", k)
	}

}

func main() {
	slice()
	sortMap()
}

//http.ParseCookie
//http.ParseSetCookie

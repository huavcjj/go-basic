package main

import "fmt"

func main() {
	//recoverから返される値には文字列が格納される
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T:%[1]s\n", e)
		}
	}()
	panic("my error") //string:my error
}

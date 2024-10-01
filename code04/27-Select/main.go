package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// i < 5 にすると、deadlock になる
	// c1とc2にそれぞれ1回ずつしかメッセージが送られていないため、ループの後半では新しいメッセージを受信できず、待機状態が続く
	// 結果的に、何も送信されないチャネルを待ち続けることになり、デッドロックが発生する
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

//eceived one
//received two

package main

import (
	"fmt"
	"time"
)

// ゴルーチンの完了を待つために、ブロッキング受信を使う
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// チャネルから値を受信するまでブロック
	<-done
}

//working...done

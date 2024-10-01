package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true //ここでチャネルにtrueを送信して、無限ループを抜ける
	fmt.Println("Ticker stopped")
}

// Tick at 2024-10-01 21:52:41.022413875 +0900 JST m=+0.500078126
// Tick at 2024-10-01 21:52:41.522397167 +0900 JST m=+1.000077209
// Tick at 2024-10-01 21:52:42.022384167 +0900 JST m=+1.500079709
// Ticker stopped

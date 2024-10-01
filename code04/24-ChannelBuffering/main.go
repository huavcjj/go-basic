package main

import "fmt"

func main() {

	// デフォルトでは、チャネルはバッファリングされないので、受信 (<- chan) の準備ができていないと送信 (chan <-) できない
	// しかし、バッファリングされたチャネルは、対応する受信側がいなくても決められた量までなら値を送信できる
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel
}

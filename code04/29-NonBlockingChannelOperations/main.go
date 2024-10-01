package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: //messagesチャネルには、バッファもなければ受ける側もいないため、msgを送信することはできない
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	msg2 := make(chan string)
	go func() {
		fmt.Println("Sending message...")
		msg2 <- "Hello"
		fmt.Println("Message sent!")
	}()

	fmt.Println("Waiting for message...")
	msg3 := <-msg2 // 受信者がないとこの時点でブロックされる
	fmt.Println("Received message:", msg3)
}

// no message received
// no message sent
// no activity

// Sending message...
// Waiting for message...
// Message sent!
// Received message: Hello

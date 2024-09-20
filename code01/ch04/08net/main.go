package main

import (
	"fmt"
	"log"
	"net"
)

func handlerData(conn net.Conn) {
	defer conn.Close()
	var b [1024]byte
	for {
		n, err := conn.Read(b[:])
		if err != nil {
			break
		}
		fmt.Print(string(b[:n]))
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlerData(conn)
	}
}

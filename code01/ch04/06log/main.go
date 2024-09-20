package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("output.dat", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started")
}

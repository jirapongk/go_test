package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jirapongk/go-test/toledo"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s\n", id)

	toledo.SayHelloToledo()
	toledo.SayHelloTest()
	toledo.SayHelloToledoPrivate()
}

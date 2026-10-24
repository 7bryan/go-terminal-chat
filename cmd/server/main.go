package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/7bryan/go-terminal-chat/internal/chat"
)

func main() {
	portFlag := flag.String("port", "8989", "TCP port number for server")
	flag.Parse()

	port := *portFlag
	if flag.NArg() > 0 {
		port = flag.Arg(0)
	}

	fmt.Printf("Starting server on port: %s\n", port)

	room := chat.NewRoom()
	go room.Run()

	if err := chat.StartServer(port, room); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}

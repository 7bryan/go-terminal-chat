package chat

import (
	"fmt"
	"log"
	"net"
)

func StartServer(port string, room *Room) error {
	addr := net.JoinHostPort("", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Failed to start server on %s: %w", addr, err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s...", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		client := &Client{
			conn: conn,
			room: room,
		}

		go client.Start()
	}
}

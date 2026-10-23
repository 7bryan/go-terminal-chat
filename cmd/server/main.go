package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: please provide a port number. Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	fmt.Printf("Starting server on port: %s\n", port)
}

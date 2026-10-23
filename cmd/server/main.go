package main

import (
	"flag"
	"fmt"
	"os"
)

type Room struct {
	Name string
	Port int
}

func main() {
	flag.Parse()

	trailingArgs := flag.Args()

	if len(trailingArgs) < 1 {
		fmt.Println("Error: please provide a port number. Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := trailingArgs[0]
	fmt.Printf("Starting server on port: %s\n", port)
}

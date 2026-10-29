# GO TERMINAL CHAT

A simple, real time collaborative terminal chat application build in Go.
The server allow multiple users to connect simultaneously via TCP, choose a username and broadcast messages to the room.

## Features

- TCP Connections: Accepts multiple concurent client connections
- System Announcements: Automatically broadcasts when users join or leave the chat
- Sender identifications: Visually distinguishes the user messages ibn the chat history by appending (you) next to the user username
- Custom Port Binding: Set a custom port using command line flags (defaults to 8989)

## Getting Started

### Prerequisites

- Go
- A TCP client like [netcat](https://nmap.org/ncat) or telnet to connect to the server

### Running the Server

Start the server using the Go CLI. By default, it will listen to port 8989

```bash
go run cmd/server/main.go
```

To specify a custom port, use the --port flag or pass the port a a direct arguemet

```bash
go run cmd/server/main.go --port 3000
```

or

```bash
go run cmd/server/main.go 3000
```

### Connecting to the Chat

Once the server is running, connect to it using netcat

```bash
nc localhost 8989
```

or

```bash
netcat localhost 8989
```

## Project Structure

```
terminal-chat/
├── cmd/
│   └── server/
│       └── main.go       # Entry point: parses port flags and starts the server
├── internal/
│   ├── chat/
│   │   ├── server.go     # TCP listener and connection acceptance loop
│   │   ├── client.go     # Struct and methods for individual user state
│   │   └── room.go       # Broadcaster logic and active connection registry
├── go.mod
└── README.md
```

# GO-TERMINAL-CHAT

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

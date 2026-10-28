package chat

import (
	"fmt"
	"sync"
)

type Message struct {
	Sender   *Client
	Time     string
	Username string
	Text     string
	System   string
}

type Room struct {
	clients   map[*Client]bool
	mu        sync.RWMutex
	broadcast chan Message
}

func NewRoom() *Room {
	return &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan Message, 256),
	}
}

func (r *Room) Run() {
	for msg := range r.broadcast {
		r.mu.RLock()

		activeClients := make([]*Client, 0, len(r.clients))
		for client := range r.clients {
			activeClients = append(activeClients, client)
		}
		r.mu.RUnlock()

		for _, client := range activeClients {
			var output string

			if msg.System != "" {
				output = msg.System
			} else if client == msg.Sender {
				output = fmt.Sprintf("[%s] %s (you): %s\n", msg.Time, msg.Username, msg.Text)
			} else {
				output = fmt.Sprintf("[%s] %s: %s\n", msg.Time, msg.Username, msg.Text)
			}
			client.conn.Write([]byte(output))
		}
	}
}

func (r *Room) AddClient(c *Client) {
	r.mu.Lock()
	r.clients[c] = true
	r.mu.Unlock()
}

func (r *Room) RemoveClient(c *Client) {
	r.mu.Lock()
	delete(r.clients, c)
	r.mu.Unlock()
}

package chat

import (
	"sync"
)

type Room struct {
	clients   map[*Client]bool
	mu        sync.RWMutex
	broadcast chan string
}

func NewRoom() *Room {
	return &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan string, 256),
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
			client.conn.Write([]byte(msg))
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

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
		broadcast: make(chan string),
	}
}

func (r *Room) Run() {
	for {
		msg := <-r.broadcast

		r.mu.RLock()

		for client := range r.clients {
			client.conn.Write([]byte(msg))
		}

		r.mu.RUnlock()
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

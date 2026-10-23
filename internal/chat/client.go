package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Client struct {
	conn     net.Conn
	username string
	room     *Room
}

func (c *Client) Start() error {
	defer func() {
		c.room.RemoveClient(c)
		c.conn.Close()
	}()

	prompt := "Choose a username: "
	if _, err := c.conn.Write([]byte(prompt)); err != nil {
		log.Printf("Failed to send prompt: %v\n", err)
		return err
	}

	scanner := bufio.NewScanner(c.conn)
	if scanner.Scan() {
		c.username = strings.TrimSpace(scanner.Text())
	}

	if err := scanner.Err(); err != nil || c.username == "" {
		log.Printf("Error reading username or empty input from %s: %v\n", c.conn.RemoteAddr(), err)
		return err
	}

	c.room.AddClient(c)
	announcement := fmt.Sprintf("=== %s has joined ===\n", c.username)
	c.room.broadcast <- announcement

	timeFormat := "15:04:05"
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		now := time.Now().Format(timeFormat)
		msg := fmt.Sprintf("[%s] %s: %s\n", now, c.username, text)
		c.room.broadcast <- msg
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error for client %s: %v", c.username, err)
		return err
	}

	departure := fmt.Sprintf("=== %s has left ===\n", c.username)
	c.room.broadcast <- departure

	return nil
}

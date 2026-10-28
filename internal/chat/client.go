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
		if c.username != "" {
			c.room.RemoveClient(c)
			departure := fmt.Sprintf("=== %s has left ===\n", c.username)
			c.room.broadcast <- Message{System: departure}
		}
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
		log.Printf("Error reading username from %s: %v\n", c.conn.RemoteAddr(), err)
		return err
	}

	c.room.AddClient(c)
	announcement := fmt.Sprintf("=== %s has joined ===\n", c.username)
	c.room.broadcast <- Message{System: announcement}

	timeFormat := "15:04:05"
	for scanner.Scan() {
		c.conn.Write([]byte("\033[1A\033[2K\r"))

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		now := time.Now().Format(timeFormat)
		c.room.broadcast <- Message{
			Sender:   c,
			Time:     now,
			Username: c.username,
			Text:     text,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error for client %s: %v", c.username, err)
		return err
	}

	return nil
}

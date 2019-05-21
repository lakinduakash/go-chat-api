package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	Pool     *Pool
	Nickname string
}

type Message struct {
	Type int         `json:"type"`
	Body MessageBody `json:"body"`
}

type MessageBody struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Fatal("Err when closing connection")
		}
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var body MessageBody
		if err := json.Unmarshal([]byte(p), &body); err != nil {
			return
		}
		body.From = c.ID
		message := Message{Type: 1, Body: body}
		c.Pool.Broadcast <- &message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

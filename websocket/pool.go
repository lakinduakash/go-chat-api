package websocket

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Broadcast  chan *Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan *Message),
	}

}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			client.ID = uuid.NewV4().String()
			pool.Clients[client.ID] = client
			fmt.Println("New user", client.ID)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for _, client2 := range pool.Clients {
				fmt.Println()
				if err := client2.Conn.WriteJSON(Message{Type: 2, Body: MessageBody{Message: client.ID}}); err != nil {
					log.Fatal("Error on write")
					continue
				}

			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client.ID)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for _, client := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 3, Body: MessageBody{Message: client.ID}}); err != nil {
					log.Fatal("Error on write")
					continue
				}
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")

			if message.Body.To != "" {
				if client, ok := pool.Clients[message.Body.To]; ok == true {
					if err := client.Conn.WriteJSON(*message); err != nil {
						fmt.Println(err)
					}
				}
			} else {
				for _, client := range pool.Clients {

					if err := client.Conn.WriteJSON(*message); err != nil {
						fmt.Println(err)
						continue
					}

				}
			}

		}
	}
}

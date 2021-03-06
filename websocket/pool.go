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

var CBR = newClientBCastChannel(2)
var CBU = newClientBCastChannel(2)
var MB = newMessageBCastChannel(2)

func (pool *Pool) GetClients() map[string]*Client {
	return pool.Clients
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			client.ID = uuid.NewV4().String()
			pool.Clients[client.ID] = client
			//fmt.Println("New user", client.ID)
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			CBR.broadCast(*client)

			for _, client2 := range pool.Clients {
				if err := client2.Conn.WriteJSON(Message{Type: 2, Body: MessageBody{Message: client.ID}}); err != nil {
					log.Fatal("Error on write")
					continue
				}

			}

			for _, client2 := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 2, Body: MessageBody{Message: client2.ID, Nickname: client2.Nickname}}); err != nil {
					log.Fatal("Error on write")
					continue
				}

			}

			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client.ID)
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			CBU.broadCast(*client)
			for _, client := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 3, Body: MessageBody{Message: client.ID}}); err != nil {
					log.Fatal("Error on write")
					continue
				}
			}
			break
		case message := <-pool.Broadcast:

			MB.broadCast(*message)
			if message.Body.To != "" {
				//fmt.Println("Sending message to", message.Body.To)

				if client, ok := pool.Clients[message.Body.To]; ok == true {
					if err := client.Conn.WriteJSON(*message); err != nil {
						fmt.Println(err)
					}
				}
			} else if message.Type == 4 {

				for _, client := range pool.Clients {

					if err := client.Conn.WriteJSON(*message); err != nil {
						fmt.Println(err)
						continue
					}

				}
			} else {
				//fmt.Println("Sending message to all clients in Pool")
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

package main

import (
	"fmt"
	chat_api "github.com/lakinduakash/go-chat-api"
)

func main() {
	go func() {
		chat_api.StartSever(":28960", "/ws")
	}()

	a := chat_api.ListenClientAddChanges()
	b := chat_api.ListenClientRemoveChanges()
	c := chat_api.ListenMessageChanges()

	for {
		select {
		case c := <-a:
			fmt.Println("New user connected ", c.ID)

		case c := <-b:
			fmt.Println("User removed ", c.ID)

		case c := <-c:
			fmt.Println("New message ", c)
		}

	}
}

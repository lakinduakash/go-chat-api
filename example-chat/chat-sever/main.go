package main

import chat_api "github.com/lakinduakash/go-chat-api"

func main() {
	chat_api.StartSever(":28960", "/ws")
}

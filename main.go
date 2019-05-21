package main

import (
	"fmt"
	"github.com/lakinduakash/go-chat-api/websocket"
	"net/http"
)

const address = ":28960"

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		if _, err := fmt.Fprintf(w, "%+v\n", err); err != nil {
			return
		}
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App")
	setupRoutes()
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Cannot serve on port ", address)
		return
	}
}

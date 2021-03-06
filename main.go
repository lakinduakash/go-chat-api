package go_chat_api

import (
	"fmt"
	"github.com/lakinduakash/go-chat-api/websocket"
	"net/http"
)

var pool *websocket.Pool

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

func setupRoutes(path string) {
	pool = websocket.NewPool()
	go pool.Start()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

//port must have ":number" format
// Ex: port:=":8080"
//path is a string,which should start with "/"
func StartSever(port string, path string) {
	fmt.Println("Distributed Chat App")
	setupRoutes(path)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Cannot serve on port ", port)
		return
	}
}

//start with TLS support. Provide privatekey and cert path. Then you need to usee "wss" protocole in client
func StartSeverTLS(port string, path string, privatekey string, cert string) {
	fmt.Println("Distributed Chat App")
	setupRoutes(path)
	if err := http.ListenAndServeTLS(port, cert, privatekey, nil); err != nil {
		fmt.Println("Cannot serve on port ", port)
		return
	}
}

//Get connected client list to sever
func GetClients() map[string]*websocket.Client {
	return pool.GetClients()
}

//This function will notify when new user is registered to chat server
func ListenClientAddChanges() chan websocket.Client {
	c := make(chan websocket.Client)
	websocket.CBR.AddWorker(c)
	return c
}

//This function will notify when a user is unregiterd from chat server
func ListenClientRemoveChanges() chan websocket.Client {
	c := make(chan websocket.Client)
	websocket.CBU.AddWorker(c)
	return c
}

//This function will notify when new message is arrived to chat server
func ListenMessageChanges() chan websocket.Message {
	c := make(chan websocket.Message)
	websocket.MB.AddWorker(c)
	return c
}

# go-chat-api

[![GoDoc](https://godoc.org/github.com/lakinduakash/go-chat-api?status.svg)](https://godoc.org/github.com/lakinduakash/go-chat-api)

Required Go version >=1.11

To use module

`go get github.com/lakinduakash/go-chat-api`

You can start server on given port and path like below.
Then you can listen client registering,removing and message arriving event.
Starting sever is blocking operation. If you need to listen changes,
Use goroutine as below. 

```go
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

```

If you need to get list of clients registered currently, use `chat_api.GetClients()`.
It will return map of clients and keys which are UUIDs.

To run example go server

```bash
git clone http://github.com/lakinduakash/go-chat-api
cd go-chat-api/example-chat/chat-sever
go run main.go
```
Start Angular client

```bash
cd go-chat-api/example-chat/chat-ui
npm install
ng serve
```

package websocket

import "sync"

var mutex1 sync.Mutex
var mutex2 sync.Mutex

type MessageBCastChannel struct {
	clientList []chan Message
}

type ClientBCastChannel struct {
	clientList []chan Client
}

func NewMessageBCastChannel(size int) *MessageBCastChannel {
	return &MessageBCastChannel{clientList: make([]chan Message, size)}
}

func NewClientBCastChannel(size int) *ClientBCastChannel {
	return &ClientBCastChannel{clientList: make([]chan Client, size)}
}

func (bc *MessageBCastChannel) AddWorker(c chan Message) {
	mutex1.Lock()
	bc.clientList = append(bc.clientList, c)
	mutex1.Unlock()
}

func (bc *ClientBCastChannel) AddWorker(c chan Client) {
	mutex2.Lock()
	bc.clientList = append(bc.clientList, c)
	mutex2.Unlock()
}

func (bc *MessageBCastChannel) BroadCast(data Message) {
	mutex1.Lock()
	for _, v := range bc.clientList {
		if v != nil {
			go func() {
				v <- data
			}()
		}
	}
	mutex1.Unlock()
}

func (bc *ClientBCastChannel) BroadCast(data Client) {
	mutex2.Lock()
	for _, v := range bc.clientList {
		if v != nil {
			go func() {
				v <- data
			}()
		}
	}
	mutex2.Unlock()
}

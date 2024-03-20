package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Room struct {
	Name    string
	clients map[*Client]bool
	lock    sync.Mutex
}

func (r *Room) Join(client *Client) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.clients[client] = true

	client.lock.Lock()
	client.rooms[r.Name] = r
	client.lock.Unlock()

	r.Broadcast([]byte(fmt.Sprintf("User %s joined room %s", client.conn.RemoteAddr(), r.Name)))
}

func (r *Room) Leave(client *Client) {
	r.lock.Lock()
	defer r.lock.Unlock()
	err := client.conn.Close()
	if err != nil {
		log.Println(err)
	}

	client.lock.Lock()
	delete(client.rooms, r.Name)
	client.lock.Unlock()

	delete(r.clients, client)
}

func (r *Room) Broadcast(msg []byte) {

	for client := range r.clients {
		err := client.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println(err)
			//potential problem with client connection
			r.Leave(client)
		}

	}
}

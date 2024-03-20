package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Room struct {
	Name    string
	clients map[*Client]bool
}

func (r *Room) Join(client *Client) {

	r.clients[client] = true

	client.rooms[r.Name] = r

	r.Broadcast([]byte(fmt.Sprintf("User %s joined room %s", client.conn.RemoteAddr(), r.Name)))
}

func (r *Room) Leave(client *Client) {
	delete(r.clients, client)
}

func (r *Room) Broadcast(msg []byte) {
	for client := range r.clients {
		_ = client.conn.WriteMessage(websocket.TextMessage, msg)
	}
}

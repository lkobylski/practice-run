package main

import "github.com/gorilla/websocket"

type Client struct {
	conn  *websocket.Conn
	rooms map[string]*Room
}

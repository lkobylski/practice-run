package main

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	conn  *websocket.Conn
	rooms map[string]*Room
	lock  sync.Mutex
}

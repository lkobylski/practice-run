package main

import (
	"errors"
	"sync"
)

type Chat struct {
	rooms map[string]*Room
	lock  sync.Mutex
}

func (c *Chat) CreateRoom(name string) (*Room, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.rooms[name]; !ok {
		c.rooms[name] = &Room{
			Name:    name,
			clients: make(map[*Client]bool),
		}

		//go c.rooms[name].handleMessages()

		return c.rooms[name], nil
	}
	return nil, errors.New("room already exists")
}

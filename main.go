package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var chat = &Chat{
	rooms: make(map[string]*Room),
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req")
		serveWs(w, r)
	})

	server := &http.Server{Addr: ":8080", Handler: nil}

	go func() {
		log.Println("Server Started on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("\nShutting down gracefully...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server Exited Properly")

}

func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	log.Println("connected")
	if err != nil {
		log.Println(err)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	client := &Client{
		conn:  conn,
		rooms: make(map[string]*Room),
	}

	for {
		messageType, raw, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		if messageType != websocket.TextMessage {
			log.Println("invalid message type")
			continue
		}

		err = handleMessage(client, raw)
		if err != nil {
			log.Printf("error handling message: %s", err)
			continue
		}
	}

	log.Println("User has been disconnected")
	for _, room := range client.rooms {
		room.Leave(client)
	}
}

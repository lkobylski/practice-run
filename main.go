package main

import (
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req")
		serveWs(w, r)
	})
	fmt.Println("server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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
			return
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

		//// TODO: bug
		//err = chat.CreateRoom(string(raw))
		//if err != nil {
		//	if err := conn.WriteMessage(websocket.TextMessage, []byte(err.Error())); err != nil {
		//		log.Println(err)
		//		return
		//	}
		//	continue
		//}

		//resMessage := fmt.Sprintf("created room with name %s", string(raw))
		//if err := conn.WriteMessage(websocket.TextMessage, []byte(resMessage)); err != nil {
		//	log.Println(err)
		//	return
		//}
	}

	//clean up

}

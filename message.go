package main

import (
	"encoding/json"
	"fmt"
)

const CMD_JOIN = "join"
const CMD_LEAVE = "leave"
const CMD_SEND_MESSAGE = "send_message"

type Message struct {
	Command string          `json:"command"`
	Payload json.RawMessage `json:"payload"`
}

type JoinRoomMessage struct {
	Room string `json:"room"`
}

type LeaveRoomMessage struct {
	Room string `json:"room"`
}

type SendMessage struct {
	Content string `json:"content"`
	Room    string `json:"room"`
}

func handleMessage(client *Client, message []byte) error {
	var m Message
	err := json.Unmarshal(message, &m)
	if err != nil {
		return err
	}

	switch m.Command {
	case CMD_JOIN:
		return handleJoin(client, m)
	case CMD_LEAVE:
		return handleLeave(client, m)
	case CMD_SEND_MESSAGE:
		return handleSendMessage(client, m)
	default:
		return fmt.Errorf("unknown command: %s", m.Command)
	}

}

func handleSendMessage(client *Client, message Message) error {
	var sendMsg SendMessage
	err := json.Unmarshal(message.Payload, &sendMsg)
	if err != nil {
		return err
	}

	room, ok := chat.rooms[sendMsg.Room]
	if !ok {
		return fmt.Errorf("room %s does not exist", sendMsg.Room)
	}

	room.Broadcast([]byte(sendMsg.Content))

	return nil

}

func handleLeave(client *Client, message Message) error {
	var leaveMsg LeaveRoomMessage
	err := json.Unmarshal(message.Payload, &leaveMsg)
	if err != nil {
		return err
	}

	room, ok := chat.rooms[leaveMsg.Room]
	if !ok {
		return fmt.Errorf("room %s does not exist", leaveMsg.Room)
	}

	room.Leave(client)

	return nil
}

func handleJoin(client *Client, message Message) error {
	var joinMsg JoinRoomMessage
	err := json.Unmarshal(message.Payload, &joinMsg)
	if err != nil {
		return err
	}

	room, ok := chat.rooms[joinMsg.Room]
	if !ok {
		room, err = chat.CreateRoom(joinMsg.Room)
		if err != nil {
			return err
		}
	}

	room.Join(client)

	return nil

}

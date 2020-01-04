package main

import "fmt"

// Room represents a chat room
type Room struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewRoom makes room easily
func NewRoom() *Room {
	return &Room{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start constantly listens for anything passed to any of our Pool’s channels and then,
// if anything is received into one of these channels, it’ll act accordingly.
func (room *Room) Start() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
			fmt.Println("There are", len(room.Clients), "person")
			for client, _ := range room.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-room.Unregister:
			delete(room.Clients, client)
			fmt.Println("There are", len(room.Clients), "person")
			for client, _ := range room.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-room.Broadcast:
			fmt.Println("Sending message")
			for client, _ := range room.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

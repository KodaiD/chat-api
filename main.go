package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveWS(room *Room, w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		Conn: conn,
		Room: room,
	}

	room.Register <- client
	client.Read()
}

func main() {
	fmt.Println("Starting chat app...")
	room := NewRoom()
	go room.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(room, w, r)
	})

	http.ListenAndServe(":8080", nil)
}

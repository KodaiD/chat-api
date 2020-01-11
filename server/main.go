package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func serveWS(room *Room, w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	// icon color
	rand.Seed(time.Now().UnixNano())
	id := "#" + strconv.FormatInt(int64(math.Floor(rand.Float64()*0xFFFFFF)), 16)

	client := &Client{
		ID:   id,
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

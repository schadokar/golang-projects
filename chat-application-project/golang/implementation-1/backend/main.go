package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/schadokar/practice-stuff/chat-application-project/golang/implementation-1/backend/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket Endpoint Hit")

	// upgrade this connection to a WebSocket connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {

	pool := websocket.NewPool()
	go pool.Start()
	// make "/ws" endpoint to the `serveWs` function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

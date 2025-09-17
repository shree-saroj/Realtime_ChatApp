package main

import (
	"fmt"
	"net/http"

	"Realtime_ChatApp/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	// go websocket.Writer(ws)
	// websocket.Reader(ws)
}

func setupRoutes() {
	// map our `/ws` endpoint to the `serveWs` function
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}
func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

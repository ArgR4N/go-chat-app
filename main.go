package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		// Receive message
		mt, message, _ := ws.ReadMessage()
		log.Printf("Message received: %s", message)

		// Response message
		_ = ws.WriteMessage(mt, message)
		log.Printf("Message sent: %s", message)
	}
}

func main() {
	fmt.Println("Server Started")
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":5555", nil))
}

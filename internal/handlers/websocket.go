// internal/handlers/websocket.go
package handlers

import (
	"anonymous-messenger/internal/models"
	ws "anonymous-messenger/internal/websocket"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader configures the Gorilla WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins; adjust for production as needed
	},
}

// HandleWebSocket manages WebSocket connections for the chat application
func HandleWebSocket(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	hub.RegisterClient(conn)

	defer func() {
		hub.UnregisterClient(conn)
	}()

	for {
		// Read messages from the client
		var receivedMsg struct {
			Text string `json:"text"`
		}
		err := conn.ReadJSON(&receivedMsg)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			break
		}

		// Save the message to the database
		msg, err := models.CreateMessage(receivedMsg.Text)
		if err != nil {
			log.Printf("Error saving message to the database: %v", err)
			continue
		}

		// Broadcast the message to all clients except the sender
		data, _ := json.Marshal(&msg)
		hub.BroadcastMessage(ws.ClientMessage{
			Message: string(data),
			Sender:  conn,
		})
	}
}

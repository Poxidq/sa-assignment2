// internal/websocket/hub.go
package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan ClientMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

type ClientMessage struct {
	Message string
	Sender  *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan ClientMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				if client != message.Sender {
					err := client.WriteMessage(websocket.TextMessage, []byte(message.Message))
					if err != nil {
						log.Printf("Error writing message to client: %v", err)
						client.Close()
						delete(h.clients, client)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) RegisterClient(conn *websocket.Conn) {
	h.register <- conn
}

func (h *Hub) UnregisterClient(conn *websocket.Conn) {
	h.unregister <- conn
}

func (h *Hub) BroadcastMessage(msg ClientMessage) {
	h.broadcast <- msg
}

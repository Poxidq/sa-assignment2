// internal/server/server.go
package server

import (
	"anonymous-messenger/internal/config"
	"anonymous-messenger/internal/database"
	"anonymous-messenger/internal/handlers"
	"anonymous-messenger/internal/models"
	"anonymous-messenger/internal/websocket"
	"log"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

// Start initializes the WebSocket hub, sets up routes, and starts the HTTP server
func (s *Server) Start() {
	// Initialize WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// Setup HTTP handlers
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWebSocket(hub, w, r)
	})

	http.HandleFunc("/messages", handlers.GetMessagesHandler)           // Fetch all messages
	http.HandleFunc("/messages/count", handlers.GetMessageCountHandler) // Fetch message count
	http.HandleFunc("/", handlers.GetIndexHandler)                      // Serve the index.html page

	// Start the server
	log.Printf("Starting server on port %s...", s.cfg.Port)
	if err := http.ListenAndServe(":"+s.cfg.Port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func Init(cfg *config.Config) {
	// Auto-migrate the Message model
	database.InitDB(cfg.DatabaseURL)
	err := database.DB.AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatalf("Init failed: %v", err)
	}
}

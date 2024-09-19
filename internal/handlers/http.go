// internal/handlers/http.go
package handlers

import (
	"anonymous-messenger/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

// GetMessagesHandler handles the GET request to retrieve all messages
func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := models.GetAllMessages()
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		log.Printf("Error retrieving messages: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
		log.Printf("Error encoding messages: %v", err)
	}
}

// GetMessageCountHandler handles the GET request to retrieve the message count
func GetMessageCountHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := models.GetAllMessages()
	if err != nil {
		http.Error(w, "Failed to retrieve message count", http.StatusInternalServerError)
		log.Printf("Error retrieving message count: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	count := map[string]int{"message_count": len(messages)}
	if err := json.NewEncoder(w).Encode(count); err != nil {
		http.Error(w, "Failed to encode message count", http.StatusInternalServerError)
		log.Printf("Error encoding message count: %v", err)
	}
}

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve index.html from the public directory
	path := filepath.Join("public", "index.html")
	http.ServeFile(w, r, path)
}

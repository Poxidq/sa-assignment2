// internal/models/message.go
package models

import (
	"anonymous-messenger/internal/database"
	"time"
)

// Message represents a chat message
type Message struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

// CreateMessage saves a new message to the database
func CreateMessage(text string) (*Message, error) {
	msg := &Message{
		Text:      text,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := database.DB.Create(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

// GetAllMessages retrieves all messages from the database
func GetAllMessages() ([]Message, error) {
	var messages []Message
	if err := database.DB.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

// internal/database/db.go
package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and migrates the message model
func InitDB(databaseURL string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connected and migrated successfully.")
	return nil
}

// cmd/server/main.go
package main

import (
	"anonymous-messenger/internal/config"
	"anonymous-messenger/internal/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	server.Init(cfg)

	// Start the server
	srv := server.NewServer(cfg)
	srv.Start()
}

# Anonymous Messenger 

## Task Overview

This assignment involves the development of a client-server application with anonymous messaging capabilities. The goal is to create a chat application where users can send messages anonymously. The system should also be able to respond to queries about the total number of messages sent.

---

## Project Structure

```
anonymous-messenger/
│
├── cmd/
│   └── main.go                   # Entry point for the server
│
├── internal/
│   ├── config/
│   │   └── config.go             # Configuration settings
│   │
│   ├── database/
│   │   └── db.go                 # Database connection and queries with GORM
│   │
│   ├── handlers/
│   │   ├── websocket.go          # WebSocket handling
│   │   └── http.go               # HTTP endpoints handling
│   │
│   ├── models/
│   │   └── message.go            # Message struct and related methods
│   │
│   ├── server/
│   │   └── server.go             # Server setup and initialization
│   │
│   └── websocket/
│       └── hub.go                # WebSocket hub to manage clients and broadcasting
│
├── pkg/
│   └── logger/
│       └── logger.go             # Logging utility for standardized logging
│
├── public/
│   └── index.html                # Frontend page for chat interaction
│
├── go.mod                        # Go module file
└── go.sum                        # Go dependencies file
```

---

## Getting Started

### Prerequisites

- Go 1.21+
- SQLite

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-repo/anonymous-messenger.git
   cd anonymous-messenger
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the server:**

   ```bash
   go run cmd/server/main.go
   ```

### Accessing the Application

- Open your browser and navigate to `http://localhost:8080/` to access the chat interface.

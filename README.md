# Chat API 💬

A lightweight Go-based REST API for handling chat messages within a forum application. This API provides endpoints for creating, retrieving, and managing chat message statuses using SQLite as the database.

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=for-the-badge&logo=sqlite&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

</div>

<details>
<summary>📋 <b>Table of Contents</b></summary>

- [Chat API 💬](#chat-api-)
  - [✨ Features](#-features)
  - [🛠 Tech Stack](#-tech-stack)
  - [📁 Project Structure](#-project-structure)
  - [🚀 Getting Started](#-getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running with Docker](#running-with-docker)
  - [📚 API Documentation](#-api-documentation)
    - [Message Model](#message-model)
    - [Endpoints](#endpoints)
      - [Available Actions](#available-actions)
  - [💡 Usage Examples](#-usage-examples)
    - [1. Create a Chat Message](#1-create-a-chat-message)
    - [2. Get All Chats](#2-get-all-chats)
    - [3. Update Message Status (Received)](#3-update-message-status-received)
    - [4. Update Message Status (Read)](#4-update-message-status-read)
  - [🔧 Development](#-development)
    - [Building from Source](#building-from-source)
    - [Running Tests](#running-tests)
    - [Environment Variables](#environment-variables)
  - [🤝 Contributing](#-contributing)

</details>

## ✨ Features

- Create chat messages between users
- Retrieve all chat messages
- Update message delivery status (received)
- Update message read status
- SQLite database for persistent storage
- Docker support for easy deployment
- RESTful API design

## 🛠 Tech Stack

- **Language**: Go 1.20
- **Database**: SQLite3
- **Docker**: Alpine-based containerization
- **Dependencies**: 
  - [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver for Go

## 📁 Project Structure

```
chatAPI/
├── config/              # Configuration constants
│   └── constants.go
├── databases/           # Database files and SQL scripts
│   └── sqlRequests/
│       ├── createTable.sql
│       └── insertNewChat.sql
├── internals/           # Internal application logic
│   ├── dbManager/       # Database initialization
│   │   └── initDB.go
│   ├── handlers/        # HTTP request handlers
│   │   ├── createHandler.go
│   │   ├── getHandler.go
│   │   ├── mainHandler.go
│   │   └── updateHandler.go
│   └── tools/           # Utility functions
│       └── utils.go
├── models/              # Data models
│   ├── message.go
│   └── request.go
├── scripts/             # Utility scripts
│   ├── init.sh
│   └── push.sh
├── Dockerfile           # Docker configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── main.go              # Application entry point
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or higher
- GCC (required for SQLite3)
- Docker (optional, for containerized deployment)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd chatAPI
   ```

2. **Install dependencies**
   ```bash
   go get github.com/mattn/go-sqlite3
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8083`
   
   > **Note:** The database will be initialized automatically on first run.

### Running with Docker

1. **Build the Docker image**
   ```bash
   docker build -t chatapi .
   ```

2. **Run the container**
   ```bash
   docker run -p 8083:8083 -v $(pwd)/databases:/app/databases chatapi
   ```

## 📚 API Documentation

### Message Model

```go
{
  "messageID": int,
  "senderID": int,
  "receiverID": int,
  "content": string,
  "statusReceived": bool,
  "statusRead": bool,
  "createAt": timestamp
}
```

### Endpoints

All endpoints use **POST** requests to `http://localhost:8083/`

The API uses an action-based request format:

```json
{
  "action": "actionName",
  "body": { ... }
}
```

#### Available Actions

| Action | Description | Requires Body |
|--------|-------------|---------------|
| `createChat` | Create a new chat message | Yes |
| `getChats` | Retrieve all chat messages | No |
| `updateStatusReceived` | Mark a message as received | Yes |
| `updateStatusRead` | Mark a message as read | Yes |

## 💡 Usage Examples

### 1. Create a Chat Message

**Request:**
```bash
curl -X POST http://localhost:8083/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "createChat",
    "body": {
      "senderID": 1,
      "receiverID": 2,
      "content": "Hello Casca, it'\''s Guts"
    }
  }'
```

**Response:**
```
Status: 201 Created
Body: "New chat created"
```

### 2. Get All Chats

**Request:**
```bash
curl -X POST http://localhost:8083/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getChats"
  }'
```

**Response:**
```
Status: 200 OK
Body: [
  {
    "messageID": 1,
    "senderID": 1,
    "receiverID": 2,
    "content": "Hello Casca, it's Guts",
    "statusReceived": false,
    "statusRead": false,
    "createAt": "2025-10-14T10:30:00Z"
  }
]
```

### 3. Update Message Status (Received)

**Request:**
```bash
curl -X POST http://localhost:8083/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "updateStatusReceived",
    "body": {
      "id": 1
    }
  }'
```

**Response:**
```
Status: 200 OK
Body: "statusReceived is updated"
```

### 4. Update Message Status (Read)

**Request:**
```bash
curl -X POST http://localhost:8083/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "updateStatusRead",
    "body": {
      "id": 1
    }
  }'
```

**Response:**
```
Status: 200 OK
Body: "statusRead is updated"
```

## 🔧 Development

### Building from Source

```bash
# Build the binary
go build -o chatapi-server

# Run the binary
./chatapi-server
```

### Running Tests

```bash
go test ./...
```

### Environment Variables

You can modify the server port in `config/constants.go`:
```go
const Port = "8083"
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

**Note:** This API is designed to work as part of a larger forum application and may require integration with authentication and user management services.

---

<div align="center">

**⭐ Star this repository if you found it helpful! ⭐**

Made with ❤️ from 🇸🇳

</div>
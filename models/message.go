package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Message struct {
	Id             int       `json:"messageID"`
	SenderId       int       `json:"senderID"`
	ReceiverId     int       `json:"receiverID"`
	Content        string    `json:"content"`
	StatusReceived bool      `json:"statusReceived"`
	StatusRead     bool      `json:"statusRead"`
	CreateAt       time.Time `json:"createAT"`
}

func (message *Message) Create(db *sql.DB) error {
	message.StatusReceived = false
	message.StatusRead = false
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewChat.sql")
	if err != nil {
		return err
	}

	// This code snippet is preparing a SQL statement for execution within a transaction.
	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		message.SenderId,
		message.ReceiverId,
		message.Content,
		message.StatusReceived,
		message.StatusRead,
		time.Now().Format(time.RFC3339),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}

	return err
}

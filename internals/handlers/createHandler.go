package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"chat/internals/tools"
	md "chat/models"
)

func createMessage(w http.ResponseWriter, message md.Message, db *sql.DB) {
	if err := message.Create(db); err != nil {
		http.Error(w, "Error while creating chat : "+err.Error(), http.StatusInternalServerError)
		return
	}

	query := `
		SELECT id, senderId, receiverId, content, statusReceived, statusRead, createdAt
		FROM chats
		ORDER BY id DESC
		LIMIT 1;
	`

	var msg md.Message
	err := db.QueryRow(query).Scan(&msg.Id, &msg.SenderId, &msg.ReceiverId, &msg.Content, &msg.StatusReceived, &msg.StatusRead, &msg.CreateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No message find")
		} else {
			http.Error(w, "Error while getting message data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	tools.WriteResponse(w, msg, http.StatusOK)
}

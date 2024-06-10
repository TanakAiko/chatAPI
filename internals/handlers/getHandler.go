package handlers

import (
	"chat/internals/tools"
	md "chat/models"
	"database/sql"
	"net/http"
)

func getChats(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT id, senderId, receiverId, content, statusReceived, statusRead, createdAt FROM chats ORDER BY createdAt ASC")
	if err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	chats := []md.Message{}
	for rows.Next() {
		var message md.Message
		if err := rows.Scan(&message.Id, &message.SenderId, &message.ReceiverId, &message.Content, &message.StatusReceived, &message.StatusRead, &message.CreateAt); err != nil {
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		chats = append(chats, message)
	}
	tools.WriteResponse(w, chats, http.StatusOK)
}

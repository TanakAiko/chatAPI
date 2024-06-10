package handlers

import (
	"database/sql"
	"net/http"

	"chat/internals/tools"
	md "chat/models"
)

func createMessage(w http.ResponseWriter, message md.Message, db *sql.DB) {
	if err := message.Create(db); err != nil {
		http.Error(w, "Error while creating chat : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New chat created", http.StatusCreated)
}

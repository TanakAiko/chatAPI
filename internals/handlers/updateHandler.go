package handlers

import (
	"chat/internals/tools"
	md "chat/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func updateStatus(choice string, w http.ResponseWriter, message md.Message, db *sql.DB) {
	sqlQuery := fmt.Sprintf("UPDATE chats SET %s = TRUE WHERE id = %d", choice, message.Id)

	result, err := db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, "Error while updating status : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No chat found with ID: "+strconv.Itoa(message.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, choice+" is updated", http.StatusOK)
}

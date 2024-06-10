package handlers

import (
	dbManager "chat/internals/dbManager"
	md "chat/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbManager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		http.Error(w, "database can't be opened", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var req md.Request
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("req.Action : ", req.Action)

	switch req.Action {
	case "createChat":
		createMessage(w, req.Body, db)
	case "getChats":
		getChats(w, db)
	case "updateStatusReceived":
		updateStatus("statusReceived", w, req.Body, db)
	case "updateStatusRead":
		updateStatus("statusRead", w, req.Body, db)
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

}

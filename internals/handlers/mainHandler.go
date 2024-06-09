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
	case "a":
		fmt.Println("something")
	case "b":
		fmt.Println("something")
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

}

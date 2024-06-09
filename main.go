package main

import (
	hd "chat/internals/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "8083"
	http.HandleFunc("/", hd.MainHandler)
	log.Printf("Server (portAPI) started at http://localhost:%v\n", port)
	http.ListenAndServe(":"+port, nil)
}

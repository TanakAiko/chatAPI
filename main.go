package main

import (
	conf "chat/config"
	hd "chat/internals/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", hd.MainHandler)
	//log.Printf("Server (chatAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}

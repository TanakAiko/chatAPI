package models

import "time"

type Message struct {
	Id         int
	SenderId   int    `json:"senderID"`
	RecieverId int    `json:"recieverID"`
	Content    string `json:"content"`
	CreateAt   time.Time
}

package models

type Request struct {
	Action string  `json:"action"`
	Body   Message `json:"body"`
}

package dto

type Message struct {
	Event   string `json:"event"`
	Payload any    `json:"payload"`
}

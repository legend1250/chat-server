package main
// Type == 1 : join a Room

type Message struct {
	Type int `json:"type"`
	RoomID string `json:"roomId"`
	Message string `json:"message"`
}
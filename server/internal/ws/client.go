package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	Username string `json:"Username"`
	ID       string `json:"Id"`
	RoomID   string `json:"RoomId"`
}

type Message struct {
	Username string `json:"Username"`
	RoomID   string `json:"RoomID"`
	Content  string `json:"Content"`
}

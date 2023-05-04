package ws

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	conn     *websocket.Conn
	Message  chan *Message
	Username string `json:"Username"`
	ID       string `json:"Id"`
	RoomID   int    `json:"RoomId"`
}

type Message struct {
	Username string `json:"Username"`
	RoomID   string `json:"RoomID"`
	Content  string `json:"Content"`
}

package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

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

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &Message{
			Content:  string(m),
			Username: c.Username,
			RoomID:   c.RoomID,
		}
		h.Broadcast <- msg
	}
}

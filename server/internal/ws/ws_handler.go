package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

type CreateUserReq struct {
	Name string `json:"name"`
	Id   string `json:"Id"`
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//we will now store it in the rooms

	h.hub.Rooms[req.Id] = &Room{
		ID:      req.Id,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//now we have a URL which will contain roomid, username, clientid. we have to extract them

	RoomID := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		Username: username,
		RoomID:   RoomID,
		ID:       clientId,
	}

	m := &Message{
		Content:  "New User has joined the room",
		RoomID:   RoomID,
		Username: username,
	}

	h.hub.Register <- cl

	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)

}

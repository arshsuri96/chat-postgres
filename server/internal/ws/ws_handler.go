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

//join room//upgrade the connection//extract by using c.params//

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//handler get room return gin

func (h *Handler) GetRoom(c *gin.Context) {
	rooms := make([]RoomRes, 0)
	for _, r := range rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	c.JSON(http.StatusOK, rooms)
}

type ClientsRes struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {

	var clients []ClientsRes

	RoomID := c.Param("roomId")
	if _, ok := h.hub.Rooms[RoomID]; ok {
		clients = make([]ClientsRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[RoomID].Clients {
		clients = append(clients, ClientsRes{
			ID:       c.ID,
			UserName: c.Username,
		})
	}
	c.JSON(http.StatusOK, clients)
}

package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	hub *Hub
}

type CreateUserReq struct {
	Name string `json:"name"`
	Id   string `json:"Id"`
}

func NewHandler(h *Hub) *handler {
	return &handler{
		hub: h,
	}
}

func (h *handler) createUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.Id] = &Room{
		Id:      req.Id,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	c.JSON(http.StatusOK, req)
}

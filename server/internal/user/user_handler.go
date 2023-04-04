package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service
}

func NewHandler(s service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) createUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.createUser(c.Request.Context(), &u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

package router

import (
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)
	r.POST("login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/JoinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/GetRoom", wsHandler.GetRoom)
	r.GET("/ws/GetClients", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}

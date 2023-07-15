package server

import (
	"business/dm"
	"business/messages"
	"business/rooms"
	"business/users"
	"context"
	"github.com/gin-gonic/gin"
	webSocket "presenter/server/web-socket"
)

func Routes(server *gin.Engine) {
	ctx := context.Background()
	manager := webSocket.NewManager(ctx)

	server.GET("/ws", manager.ServeWS)

	// Register routes
	server.POST("/register", users.Register(manager))
	server.POST("/login", users.Login(manager))

	// Rooms routes
	server.GET("/rooms/get-all", rooms.GetAll)
	server.GET("/room/get-messages", messages.Fetch)

	// Users
	server.GET("/users/fetch", users.Fetch)

	// DM - fetch or create DM between two users
	server.POST("dm/fetch", dm.Fetch)
	server.GET("dm/messages", dm.FetchMessages)
	server.POST("dm/add-message", dm.AddMessage)
}

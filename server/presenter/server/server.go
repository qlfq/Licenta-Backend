package server

import (
	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	Routes(router)

	router.Run(":8080")
}

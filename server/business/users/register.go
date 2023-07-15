package users

import (
	data "data/users"
	"domain/users/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	web_socket "presenter/server/web-socket"
)

func Register(m *web_socket.Manager) func(*gin.Context) {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			fmt.Println("Error")
			c.IndentedJSON(http.StatusInternalServerError, "")
			return
		}

		result, err := data.CreateUser(user)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		c.IndentedJSON(http.StatusCreated, createOTP(m, result))
	}
}

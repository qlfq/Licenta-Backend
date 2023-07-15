package users

import (
	"data/users"
	"domain/users/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	web_socket "presenter/server/web-socket"
)

func Login(m *web_socket.Manager) func(*gin.Context) {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "")
			return
		}

		result, err := users.Login(user)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err)
			return
		}

		c.IndentedJSON(http.StatusOK, createOTP(m, result))
	}
}

func validatePassword(found string, inserted string) bool {
	if found == inserted {
		return true
	}

	return false
}

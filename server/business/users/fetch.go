package users

import (
	"data/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Fetch(c *gin.Context) {
	userId := c.Query("user-id")
	uintId, _ := strconv.ParseUint(userId, 10, 32)

	result, err := users.Fetch(uint(uintId))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Failed to fetch the users")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

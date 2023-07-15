package dm

import (
	"data/dm"
	"domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fetch(c *gin.Context) {
	var payload domain.DM

	if err := c.BindJSON(&payload); err != nil {
		fmt.Println("Error")
		c.IndentedJSON(http.StatusBadRequest, "Failed to bind payload")
		return
	}

	result, err := dm.Fetch(payload)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Failed to fetch dm")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
	return
}

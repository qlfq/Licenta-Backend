package dm

import (
	"data/dm"
	"domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMessage(c *gin.Context) {
	var message domain.DmMessage

	if err := c.BindJSON(&message); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Failed to bind payload")
		return
	}

	result, err := dm.AddMessage(message)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Failed to add the new message")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

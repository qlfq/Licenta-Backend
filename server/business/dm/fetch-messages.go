package dm

import (
	"data/dm"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FetchMessages(c *gin.Context) {
	userId := c.Query("dm-id")
	uintId, _ := strconv.ParseUint(userId, 10, 32)

	result, err := dm.FetchMessages(uint(uintId))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

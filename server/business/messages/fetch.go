package messages

import (
	"data/messages"
	"domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func convertString(queryParam string) (uint, error) {
	result, err := strconv.Atoi(queryParam)

	if err != nil {
		return 0, err
	}

	return uint(result), nil
}

func convertMessages(messages []domain.Message) []domain.Message {
	var result []domain.Message

	for _, message := range messages {
		message.Decrypt()

		result = append(result, message)
	}

	return result
}

func Fetch(c *gin.Context) {
	queryParam := c.Query("roomId")

	roomId, convertErr := convertString(queryParam)

	if convertErr != nil {
		fmt.Println(convertErr)
		c.IndentedJSON(http.StatusBadRequest, "RoomId cannot be converted to UINT type")
		return
	}

	result, err := messages.Fetch(roomId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Failed to fetch the messages")
		return
	}

	c.IndentedJSON(http.StatusOK, convertMessages(result))
}

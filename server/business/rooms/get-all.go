package rooms

import (
	roomsRepository "data/rooms"
	"domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func searchForUserId(userId uint, persons []domain.RoomPerson) bool {
	for _, person := range persons {
		if person.UserId == userId {
			return true
		}
	}

	return false
}

func isCreatedByUser(userId uint, roomCreatorId uint) bool {
	if userId == roomCreatorId {
		return true
	}

	return false
}

func filterResult(userId uint, fetchedRooms []domain.Room) []domain.Room {
	var rooms []domain.Room

	for _, room := range fetchedRooms {
		if room.IsOpen == true {
			rooms = append(rooms, room)
		} else if isCreatedByUser(userId, room.UserId) {
			rooms = append(rooms, room)
		} else {
			persons, err := roomsRepository.GetRoomPeople(room.Id)

			if err == nil {
				if searchForUserId(userId, persons) {
					rooms = append(rooms, room)
				}
			}
		}
	}

	return rooms
}

func GetAll(c *gin.Context) {
	userId := c.Query("user-id")
	uintId, _ := strconv.ParseUint(userId, 10, 32)

	rooms, err := roomsRepository.GetAll()

	result := filterResult(uint(uintId), rooms)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "")
	}

	c.IndentedJSON(http.StatusOK, result)
}

package rooms

import (
	"data/rooms"
	"domain"
)

func IsOpen(roomId uint) domain.Room {
	return rooms.IsOpen(roomId)
}

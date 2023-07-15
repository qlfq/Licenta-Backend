package rooms

import (
	"data"
	"data/rooms/entities"
)

func Delete(roomId uint) bool {

	db, err := data.Context()

	if err != nil {
		panic("something went wrong")
		return false
	}

	var room entities.Room

	db.Where("ID=?", roomId).First(&room)

	room.Active = false

	db.Save(&room)

	return true
}

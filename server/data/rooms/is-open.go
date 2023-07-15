package rooms

import (
	"data"
	"data/rooms/entities"
	"domain"
)

func IsOpen(roomId uint) domain.Room {
	db, err := data.Context()

	if err != nil {
		panic("something went wrong")
		return domain.Room{}
	}

	var result entities.Room

	db.Where("id=?", roomId).Where("is_open=?", 1).Find(&result)

	return entities.MapToRoom(result)
}

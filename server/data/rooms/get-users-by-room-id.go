package rooms

import (
	"data"
	"data/rooms/entities"
	"domain"
)

func GetUsersByRoomId(id uint) ([]domain.RoomPerson, error) {
	db, err := data.Context()

	if err != nil {
		return make([]domain.RoomPerson, 0), err
	}

	var result []entities.RoomPerson
	db.Where("room_id=?", id).Find(&result)

	return entities.MapToRoomPersons(result), nil
}

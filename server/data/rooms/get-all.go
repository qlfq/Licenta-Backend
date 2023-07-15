package rooms

import (
	"data"
	"data/rooms/entities"
	"domain"
)

func GetAll() ([]domain.Room, error) {
	db, err := data.Context()

	if err != nil {
		panic("something went wrong")
		return make([]domain.Room, 0), err
	}

	var result []entities.Room

	db.Where("active=1").Find(&result)

	return entities.MapToRooms(result), nil
}

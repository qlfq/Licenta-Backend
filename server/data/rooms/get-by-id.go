package rooms

import (
	"data"
	"data/rooms/entities"
	"domain"
)

func GetById(id uint) (domain.Room, error) {
	db, err := data.Context()

	if err != nil {
		return domain.Room{}, err
	}

	var result entities.Room

	db.Where("Id=?", id).First(&result)

	return entities.MapToRoom(result), nil
}

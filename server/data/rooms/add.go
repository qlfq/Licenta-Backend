package rooms

import (
	"data"
	"data/rooms/entities"
	"domain"
)

func Add(room domain.Room) (domain.Room, error) {
	db, err := data.Context()
	insert := entities.MapToERoom(room)

	if len(room.Persons) == 0 && !room.IsOpen {
		insert.IsOpen = true
	}

	if err != nil {
		panic("something went wrong")
		return domain.Room{}, err
	}

	if value := db.Create(&insert); value.Error != nil {
		return domain.Room{}, value.Error
	}

	if len(room.Persons) != 0 {
		var insertedRoom entities.Room
		db.Where("name=?", room.Name).First(&insertedRoom)

		insertRoomPersons := entities.MapToERoomPersons(room.Persons, insertedRoom)

		if value := db.Create(&insertRoomPersons); value.Error != nil {
			return domain.Room{}, value.Error
		}
	}

	var result entities.Room

	db.First(&result, "name=?", insert.Name)

	return entities.MapToRoom(result), nil
}

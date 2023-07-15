package rooms

import (
	"data/rooms"
	"domain"
)

func Add(room domain.Room) (domain.Room, error) {
	result, err := rooms.Add(room)

	if err != nil {
		return domain.Room{}, err
	}

	return result, nil
}

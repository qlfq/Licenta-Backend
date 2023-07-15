package rooms

import (
	"data"
	"domain"
)

func GetRoomPeople(roomId uint) ([]domain.RoomPerson, error) {
	db, err := data.Context()
	var result []domain.RoomPerson

	if err != nil {
		panic("Something went wrong")
		return make([]domain.RoomPerson, 0), err
	}

	db.Where("room_id=?", roomId).Find(&result)

	return result, nil
}

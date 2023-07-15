package messages

import (
	"data"
	"data/messages/model"
	"domain"
)

func Fetch(roomId uint) ([]domain.Message, error) {
	db, err := data.Context()

	if err != nil {
		panic("something went wrong")
		return make([]domain.Message, 0), err
	}

	var result []model.Message

	db.Where("room_id=?", roomId).Find(&result)

	return model.MapToMessages(result), nil
}

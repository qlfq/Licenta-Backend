package messages

import (
	"data"
	"data/messages/model"
	"domain"
)

func Add(message domain.Message) (domain.Message, error) {
	db, err := data.Context()
	insert := model.MapToEMessage(message)

	if err != nil {
		panic("something went wrong")
		return domain.Message{}, err
	}

	if value := db.Create(&insert); value.Error != nil {
		return domain.Message{}, value.Error
	}

	return domain.Message{}, nil
}

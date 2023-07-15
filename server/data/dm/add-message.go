package dm

import (
	"data"
	"data/dm/models"
	"domain"
)

func AddMessage(message domain.DmMessage) ([]domain.DmMessage, error) {
	db, err := data.Context()
	insert := models.MapToEMessage(message)

	if err != nil {
		return make([]domain.DmMessage, 0), err
	}

	if value := db.Create(&insert); value.Error != nil {
		return make([]domain.DmMessage, 0), value.Error
	}

	var result []models.DmMessage

	db.Where("dm_id=?", message.DmID).Find(&result)

	return models.MapToMessages(result), nil
}

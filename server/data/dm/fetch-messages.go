package dm

import (
	"data"
	"data/dm/models"
	"domain"
)

func FetchMessages(dmId uint) ([]domain.DmMessage, error) {
	db, err := data.Context()

	if err != nil {
		panic("something went wrong")

		return make([]domain.DmMessage, 0), err
	}

	var result []models.DmMessage

	db.Where("dm_id=?", dmId).Find(&result)

	return models.MapToMessages(result), nil
}

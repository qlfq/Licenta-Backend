package dm

import (
	"data"
	"data/dm/models"
	"domain"
)

func FetchById(id uint) (domain.DM, error) {
	db, err := data.Context()

	if err != nil {
		return domain.DM{}, err
	}

	var result models.DM

	db.Where("Id=?", id).First(&result)

	return models.MapToDM(result), nil
}

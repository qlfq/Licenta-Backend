package dm

import (
	"data"
	"data/dm/models"
	"domain"
)

func Fetch(dm domain.DM) (domain.DM, error) {
	db, err := data.Context()
	result := models.MapToEDM(dm)

	if err != nil {
		panic("something went wrong")

		return domain.DM{}, err
	}

	var firstGet models.DM
	var secondGet models.DM

	db.Where("first_user_id=?", result.FirstUserId).Where("second_user_id=?", result.SecondUserId).First(&firstGet)
	db.Where("first_user_id=?", result.SecondUserId).Where("second_user_id=?", result.FirstUserId).First(&secondGet)

	if firstGet == (models.DM{}) && secondGet == (models.DM{}) {
		db.Create(&result)

		return models.MapToDM(result), nil
	}

	if firstGet != (models.DM{}) {
		return models.MapToDM(firstGet), nil
	}

	if secondGet != (models.DM{}) {
		return models.MapToDM(secondGet), nil
	}

	return domain.DM{}, err
}

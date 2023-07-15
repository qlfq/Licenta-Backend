package users

import (
	"data"
	"data/users/model"
	"domain/users/models"
	"log"
)

func Fetch(userId uint) ([]models.User, error) {
	db, err := data.Context()

	if err != nil {
		panic("something went wrong")

		return make([]models.User, 0), err
	}

	var result []model.User

	if err := db.Where("id != ?", userId).Find(&result).Error; err != nil {
		log.Fatalln(err)
	}

	return model.MapEUsers(result), nil
}

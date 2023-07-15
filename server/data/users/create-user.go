package users

import (
	"data"
	"data/users/model"
	"domain/users/models"
)

func CreateUser(user models.User) (models.User, error) {
	db, err := data.Context()
	insert := model.MapDbUser(user)

	if err != nil {
		panic("Something went wrong")
		return models.User{}, err
	}

	if value := db.Create(&insert); value.Error != nil {
		return models.User{}, value.Error
	}

	var result model.User

	db.Where("username=?", insert.Username).First(&result)

	return model.MapUser(result), nil
}

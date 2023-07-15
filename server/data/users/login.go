package users

import (
	"data"
	"data/users/model"
	"domain/users/models"
	"errors"
	"fmt"
)

func Login(user models.User) (models.User, error) {
	db, err := data.Context()
	var result model.User

	if err != nil {
		panic("Something went wrong")
		return models.User{}, err
	}

	db.Where("username=?", user.Username).First(&result)

	password := result.Password.Raw.(string)

	fmt.Printf("password %v", password)
	fmt.Printf("passwordFromFE %v: ", user.Password)

	if user.Password == password {
		return model.MapUser(result), nil
	}

	return models.User{}, errors.New("something went wrong")
}

package model

import (
	"domain/users/models"
	gormcrypto "github.com/pkasila/gorm-crypto"
)

func MapDbUser(user models.User) User {
	return User{
		Username: user.Username,
		Password: gormcrypto.EncryptedValue{Raw: user.Password},
		Email:    gormcrypto.EncryptedValue{Raw: user.Email},
	}
}

func MapUser(user User) models.User {
	return models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email.Raw.(string),
	}
}

func MapEUsers(users []User) []models.User {
	var result []models.User

	for _, user := range users {
		result = append(result, MapUser(user))
	}

	return result
}

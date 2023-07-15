package data

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Context() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")

		return nil, err
	}

	return db, nil
}

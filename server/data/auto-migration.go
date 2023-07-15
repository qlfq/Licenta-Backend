package data

import (
	"data/dm/models"
	"data/messages/model"
	"data/rooms/entities"
	model3 "data/users/model"
	"log"
)

func AutoMigration() {
	db, err := Context()

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&model.Message{}, &entities.RoomPerson{}, &entities.Room{}, &model3.User{}, models.DM{}, models.DmMessage{}); err != nil {
		log.Fatalln(err)
	}
}

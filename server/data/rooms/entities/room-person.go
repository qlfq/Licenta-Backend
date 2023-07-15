package entities

import "gorm.io/gorm"

type RoomPerson struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserId uint
	RoomId uint
	Label  string
}

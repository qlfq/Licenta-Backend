package models

import "gorm.io/gorm"

type DM struct {
	gorm.Model
	ID           uint `gorm:"primaryKey:autoIncrement"`
	FirstUserId  uint
	SecondUserId uint
}

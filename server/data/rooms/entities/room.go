package entities

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	UserId    uint
	IsOpen    bool
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

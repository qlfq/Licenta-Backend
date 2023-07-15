package model

import (
	gormcrypto "github.com/pkasila/gorm-crypto"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username gormcrypto.EncryptedValue
	Date     time.Time
	Text     gormcrypto.EncryptedValue
	RoomID   uint
	UserId   uint
}

package models

import (
	gormcrypto "github.com/pkasila/gorm-crypto"
	"gorm.io/gorm"
	"time"
)

type DmMessage struct {
	gorm.Model
	Id       uint `gorm:"primaryKey:autoIncrement"`
	SenderId uint
	DmId     uint
	Message  gormcrypto.EncryptedValue
	Username gormcrypto.EncryptedValue
	CreateAt time.Time
}

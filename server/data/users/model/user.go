package model

import (
	"github.com/pkasila/gorm-crypto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Password gormcrypto.EncryptedValue
	Email    gormcrypto.EncryptedValue `gorm:"unique"`
}

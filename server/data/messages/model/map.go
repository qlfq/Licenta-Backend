package model

import (
	"domain"
	gormcrypto "github.com/pkasila/gorm-crypto"
	"time"
)

func MapToEMessage(message domain.Message) Message {
	return Message{
		Username: gormcrypto.EncryptedValue{Raw: message.Username},
		Date:     time.Now(),
		Text:     gormcrypto.EncryptedValue{Raw: message.Text},
		RoomID:   message.RoomId,
		UserId:   message.UserId,
	}
}

func MapToMessage(message Message) domain.Message {
	return domain.Message{
		Id:       message.ID,
		Username: message.Username.Raw.(string),
		UserId:   message.UserId,
		RoomId:   message.RoomID,
		Date:     time.Time{},
		Text:     message.Text.Raw.(string),
	}
}

func MapToMessages(messages []Message) []domain.Message {
	var result []domain.Message

	for _, message := range messages {
		result = append(result, MapToMessage(message))
	}

	return result
}

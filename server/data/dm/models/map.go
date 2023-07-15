package models

import (
	"domain"
	gormcrypto "github.com/pkasila/gorm-crypto"
)

func MapToEDM(dm domain.DM) DM {
	return DM{
		FirstUserId:  dm.FirstUserId,
		SecondUserId: dm.SecondUserId,
	}
}

func MapToDM(dm DM) domain.DM {
	return domain.DM{
		Id:           dm.ID,
		FirstUserId:  dm.FirstUserId,
		SecondUserId: dm.SecondUserId,
	}
}

func MapToMessage(message DmMessage) domain.DmMessage {
	return domain.DmMessage{
		Id:       message.Id,
		SenderId: message.SenderId,
		DmID:     message.DmId,
		Message:  message.Message.Raw.(string),
		CreateAt: message.CreateAt,
		Username: message.Username.Raw.(string),
	}
}

func MapToEMessage(message domain.DmMessage) DmMessage {
	return DmMessage{
		Id:       message.Id,
		SenderId: message.SenderId,
		DmId:     message.DmID,
		Message:  gormcrypto.EncryptedValue{Raw: message.Message},
		CreateAt: message.CreateAt,
		Username: gormcrypto.EncryptedValue{Raw: message.Username},
	}
}

func MapToMessages(messages []DmMessage) []domain.DmMessage {
	var result []domain.DmMessage

	for _, message := range messages {
		result = append(result, MapToMessage(message))
	}

	return result
}

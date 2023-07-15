package rooms

import (
	"data/rooms"
	"domain"
)

func extractUsersId(users []domain.RoomPerson) []uint {
	var usersId []uint

	for _, user := range users {
		usersId = append(usersId, user.UserId)
	}

	return usersId
}

func GetUsersByRoomId(id uint) ([]uint, error) {
	users, err := rooms.GetUsersByRoomId(id)

	if err != nil {
		return make([]uint, 0), err
	}

	return extractUsersId(users), nil
}

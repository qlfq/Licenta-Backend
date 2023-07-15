package entities

import (
	"domain"
)

func MapToERoom(room domain.Room) Room {
	return Room{
		Name:   room.Name,
		UserId: room.UserId,
		IsOpen: room.IsOpen,
		Active: true,
	}
}

func MapToRoom(room Room) domain.Room {
	return domain.Room{
		Id:     room.ID,
		Name:   room.Name,
		UserId: room.UserId,
		IsOpen: room.IsOpen,
	}
}

func MapToRooms(rooms []Room) []domain.Room {
	var result []domain.Room
	for _, v := range rooms {
		result = append(result, MapToRoom(v))
	}

	return result
}

func mapToERoomPerson(person domain.RoomPerson, roomId uint) RoomPerson {
	return RoomPerson{
		UserId: person.UserId,
		Label:  person.Label,
		RoomId: roomId,
	}
}

func mapToERoom(room Room) RoomPerson {
	return RoomPerson{
		UserId: room.UserId,
		RoomId: room.ID,
	}
}

func MapToERoomPersons(persons []domain.RoomPerson, insertedRoom Room) []RoomPerson {
	var result []RoomPerson
	result = append(result, mapToERoom(insertedRoom))

	for _, v := range persons {
		result = append(result, mapToERoomPerson(v, insertedRoom.ID))
	}

	return result
}

func mapToRoomPerson(person RoomPerson) domain.RoomPerson {
	return domain.RoomPerson{
		Id:     person.ID,
		RoomId: person.RoomId,
		UserId: person.UserId,
	}
}

func MapToRoomPersons(persons []RoomPerson) []domain.RoomPerson {
	var result []domain.RoomPerson

	for _, person := range persons {
		result = append(result, mapToRoomPerson(person))
	}

	return result
}

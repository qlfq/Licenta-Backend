package rooms

import (
	"data/rooms"
)

func Delete(roomId uint) bool {
	removed := rooms.Delete(roomId)

	return removed
}

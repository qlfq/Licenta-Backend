package dm

import (
	"data/dm"
)

func FetchById(id uint) ([]uint, error) {
	result, err := dm.FetchById(id)

	if err != nil {
		return make([]uint, 0), err
	}

	return []uint{result.FirstUserId, result.SecondUserId}, nil
}

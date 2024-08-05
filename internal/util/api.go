package util

import (
	"fmt"
	"strconv"
)

func GetIntId(idParam string) (int32, error) {
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return 0, fmt.Errorf("error parsing ID: %v", err)
	}

	return int32(id), nil
}

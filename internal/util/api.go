package util

import (
	"fmt"
	"strconv"
)

func GetIntId(paramId string) (int32, error) {
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return 0, fmt.Errorf("error parsing ID: %v", err)
	}

	return int32(id), nil
}

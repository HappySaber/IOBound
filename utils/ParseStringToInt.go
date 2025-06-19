package utils

import (
	"errors"
	"strconv"
)

func ParseStrToInt(str string) (int, error) {
	a, err := strconv.Atoi(str)

	if err != nil {
		return 0, errors.New("couldn't pasre string to int")
	}

	return a, nil
}

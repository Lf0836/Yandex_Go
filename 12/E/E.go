package main

import (
	"errors"
	"strconv"

)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	return strconv.FormatInt(int64(num), 2), nil
}
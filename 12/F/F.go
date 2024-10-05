package main

import (
	"errors"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	return num1 + num2, nil
}

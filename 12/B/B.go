package main

import "errors"

func DivideIntegers(n, m int) (float64, error) {
	if n == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(n) / float64(m), nil
}

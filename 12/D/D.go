package main

import "errors"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	factor := 1
	for i := 2; i <= n; i++ {
		factor *= i
	}
	return factor, nil
}

package main

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 4, -8},
		{7, -3, -21},
		{-4, -5, 20},
		{10, 10, 100},
	}
	for _, test := range tests {
		ans := Multiply(test.a, test.b)
		if ans != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", test.a, test.b, ans, test.expected)
		}
	}
}

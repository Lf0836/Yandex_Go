package main

import (
	"errors"
	"testing"

)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		wantErr  error
	}{

		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
		{[]byte(""), 0, nil},
	}

	for _, test := range tests {
		length, err := GetUTFLength(test.input)
		if length != test.expected {
			t.Errorf("GetUTFLength(%q) = %d, expected %d", test.input, length, test.expected)
		}
		if !errors.Is(err, test.wantErr) {
			t.Errorf("GetUTFLength(%q) returned error %v, expected %v", test.input, err, test.wantErr)
		}
	}
}

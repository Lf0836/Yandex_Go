package main

import "errors"

func GetCharacterAtPosition(str string, pos int) (rune, error) {
	r := []rune(str)
	if pos < 0 || pos >= len(r) {
		return 0, errors.New("position out of range")
	}
	return r[pos], nil
}

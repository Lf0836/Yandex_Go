package main

import "unicode"

func isLatin(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}

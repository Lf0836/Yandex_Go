package ex4

import "testing"

func isVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hll"},
		{"world", "wrld"},
		{"AEIOU", ""},
		{"gopher", "gphr"},
		{"This is a test", "Ths s  tst"},
		{"y", "y"},
		{"", ""},
		{"Go Programming", "G Prgrmmng"},
	}
	for _, test := range tests {
		result := DeleteVowels(test.input)
		if result != test.expected {
			t.Errorf("DeleteVowels(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

package check_pattern

import (
	"testing"
)

func TestCheckPattern(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected bool
	}{
		{"hello", "hlo", true},
		{"hello", "ohl", false},
		{"abcdefg", "abc", true},
		{"xaybzc", "abc", true},
		{"eraaaaaaaaaaer", "er", false},
		{"err", "er", true},
		{"eraaaaaaaaaaer", "era", false},
		{"abcdefg", "hij", true},
		{"aaaaaaa", "a", true},
		{"aaaabaa", "ab", false},
	}

	for _, tc := range testCases {
		result := checkPattern(tc.input, tc.pattern)
		if result != tc.expected {
			t.Errorf("Input: %s, Pattern: %s, Expected: %v, Got: %v", tc.input, tc.pattern, tc.expected, result)
		}
	}
}

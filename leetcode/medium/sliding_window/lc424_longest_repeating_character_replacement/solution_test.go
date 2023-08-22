package lc424_longest_repeating_character_replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output int
	}{
		// Test cases from the problem statement
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},

		// Additional test cases
		{"AAAA", 0, 4},        // All same characters, k=0
		{"ABCDE", 2, 3},       // Different characters, k=2
		{"ABBBBCCCDDE", 3, 7}, // Mix of characters, k=3
		{"", 0, 0},            // Empty string

		// Edge cases
		{"A", 0, 1}, // Single character, k=0
		{"", 5, 0},  // Empty string, k=5
	}

	for _, test := range tests {
		assert.Equal(t, test.output, characterReplacement(test.s, test.k),
			"Input: %s, k: %d", test.s, test.k)
	}
}

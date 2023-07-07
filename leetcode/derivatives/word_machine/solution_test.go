package word_machine

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		ops      string
		expected int
	}{
		{"13 DUP 4 POP 5 DUP + DUP + -", 7},
		{"5 6 + -", -1},
		{"3 DUP 5 - -", -1},
		{"1048575 1 +", -1},     // Overflow scenario: 1048575 + 1 = 1048576
		{"1048576 1 -", -1},     // Underflow scenario: 1 - 1048576 = -1048575
		{"1048575 DUP +", -1},   // Overflow scenario: 1048575 + 1048575 = 2097150
		{"0 1 -", 1},            // Underflow scenario: 1 - 0 = 1
		{"524288 DUP -", -1},    // Underflow scenario: 524288 - 524288 = 0
		{"524288 524288 +", -1}, // Overflow scenario: 524288 + 524288 = 1048576
	}

	for _, tc := range testCases {
		result := solution(tc.ops)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d for ops %q", tc.expected, result, tc.ops)
		}
	}
}

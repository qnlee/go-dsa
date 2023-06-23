package lc91_decode_ways

import (
	"testing"
)

func Test_NumDecodings_TopDown(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"11106", 2},
	}

	for _, tc := range testCases {
		result := numDecodingsTopDown(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}

func Test_NumDecodings_BottomUp(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"11106", 2},
	}
	for _, tc := range testCases {
		result := numDecodingsBottomUp(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}

func Benchmark_NumDecodings_TopDown(b *testing.B) {
	s := "11106226220" // Example input string
	for i := 0; i < b.N; i++ {
		numDecodingsTopDown(s)
	}
}

func Benchmark_NumDecodings_BottomUp(b *testing.B) {
	s := "11106226220" // Example input string
	for i := 0; i < b.N; i++ {
		numDecodingsBottomUp(s)
	}
}

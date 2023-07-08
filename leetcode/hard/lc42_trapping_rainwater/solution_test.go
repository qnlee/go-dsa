package lc42_trapping_rainwater

import (
	"fmt"
	"testing"
)

func TestTrapDP(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TrapDP(%v)", tt.height), func(t *testing.T) {
			got := trapDP(tt.height)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestTrapTwoPointer(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TrapTwoPointer(%v)", tt.height), func(t *testing.T) {
			got := trapTwoPointer(tt.height)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestTrapTwoPointerOpt(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TrapTwoPointerOpt(%v)", tt.height), func(t *testing.T) {
			got := trapTwoPointerOpt(tt.height)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkTrapDP(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for n := 0; n < b.N; n++ {
		trapDP(height)
	}
}

func BenchmarkTrapTwoPointer(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for n := 0; n < b.N; n++ {
		trapTwoPointer(height)
	}
}

func BenchmarkTrapTwoPointerOpt(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for n := 0; n < b.N; n++ {
		trapTwoPointerOpt(height)
	}
}

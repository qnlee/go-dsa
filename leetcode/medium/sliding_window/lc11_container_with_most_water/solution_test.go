package lc11_container_with_most_water

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("MaxArea(%v)", tt.height), func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

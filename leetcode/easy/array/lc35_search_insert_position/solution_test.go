package lc35_search_insert_position

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6}, 10, 4},
	}

	for idx, tc := range tests {
		result := searchInsert(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("Test case %d failed: Expected %d, but got %d", idx+1, tc.expected, result)
		}
	}
}

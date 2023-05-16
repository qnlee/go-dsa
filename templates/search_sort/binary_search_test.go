package search_sort_test

import (
	"go-dsa/templates/search_sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 9, expected: 8},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 0, expected: -1},
		{nums: []int{9, 5, 3, 4, 2, 6, 7, 8, 1, 10}, target: 5, expected: -1},
	}

	for _, c := range cases {
		got := search_sort.BinarySearch(c.nums, c.target)
		if got != c.expected {
			t.Errorf("BinarySearch(%v, %d) == %d, expected %d", c.nums, c.target, got, c.expected)
		}
	}
}

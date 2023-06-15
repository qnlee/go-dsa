package lc153_find_minimum_in_rotated_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinSoln1(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		// Test case where minimum element is in the middle
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},

		// Test case where minimum element is at the beginning
		{nums: []int{1, 2, 3, 4, 5}, expected: 1},

		// Test case where minimum element is at the end
		{nums: []int{6, 7, 8, 9, 1}, expected: 1},

		// Test case where all elements are the same
		{nums: []int{3, 3, 3, 3, 3}, expected: 3},

		// Test case with a single element
		{nums: []int{5}, expected: 5},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, findMinSoln1(tc.nums))
	}
}

func TestFindMinSoln2(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		// Test case where minimum element is in the middle
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},

		// Test case where minimum element is at the beginning
		{nums: []int{1, 2, 3, 4, 5}, expected: 1},

		// Test case where minimum element is at the end
		{nums: []int{6, 7, 8, 9, 1}, expected: 1},

		// Test case where all elements are the same
		{nums: []int{3, 3, 3, 3, 3}, expected: 3},

		// Test case with a single element
		{nums: []int{5}, expected: 5},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, findMinSoln2(tc.nums))
	}
}

func BenchmarkFindMinSoln1(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		findMinSoln1(nums)
	}
}

func BenchmarkFindMinSoln2(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		findMinSoln2(nums)
	}
}

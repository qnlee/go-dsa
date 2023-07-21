package lc673_number_of_longest_increasing_subsequence

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 3, 5, 4, 7},
			expected: 2,
		},
		{
			nums:     []int{2, 2, 2, 2, 2},
			expected: 5,
		},
		{
			nums:     []int{1, 2, 4, 3, 5, 4, 7, 2},
			expected: 3,
		},
		{
			nums:     []int{10},
			expected: 1,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for idx, tc := range testCases {
		result := findNumberOfLIS(tc.nums)
		if result != tc.expected {
			t.Errorf("Test case %d failed: Expected %d, but got %d", idx+1, tc.expected, result)
		}
	}
}

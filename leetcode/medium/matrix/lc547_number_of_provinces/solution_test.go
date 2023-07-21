package lc547_number_of_provinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	testCases := []struct {
		isConnected [][]int
		expected    int
	}{
		{
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
		{
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			expected: 1,
		},
		{
			isConnected: [][]int{
				{1, 0, 0, 1},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 1},
			},
			expected: 2,
		},
		{
			isConnected: [][]int{{1}},
			expected:    1,
		},
	}

	for idx, tc := range testCases {
		result := findCircleNum(tc.isConnected)
		if result != tc.expected {
			t.Errorf("Test case %d failed: Expected %d, but got %d", idx+1, tc.expected, result)
		}
	}
}

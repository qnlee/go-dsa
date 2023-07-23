package lc1492_kth_factor_of_n

import "testing"

func TestKthFactor(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected int
	}{
		{12, 3, 3},
		{7, 2, 7},
		{4, 4, -1},
	}

	for idx, tc := range tests {
		result := kthFactor(tc.n, tc.k)
		if result != tc.expected {
			t.Errorf("Test case %d failed: Expected %d, but got %d", idx+1, tc.expected, result)
		}
	}
}

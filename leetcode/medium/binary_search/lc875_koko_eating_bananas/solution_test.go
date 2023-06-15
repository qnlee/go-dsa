package lc875_koko_eating_bananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		piles    []int
		h        int
		expected int
	}{
		// Test case where Koko can eat all the bananas within the given hours
		{
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		// Test case where Koko needs to eat at the maximum speed to finish within the given hours
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		// Test case where Koko can finish within the given hours but doesn't need to eat at the maximum speed
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, minEatingSpeed(tc.piles, tc.h))
	}
}

package lc42_trapping_rainwater

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it
can trap after raining.

Example 1:
	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6

Example 2:
	Input: height = [4,2,0,3,2,5]
	Output: 9

Constraints:
	n == height.length
	1 <= n <= 2 * 104
	0 <= height[i] <= 105
*/

// O(n) time, O(n) space
func trapDP(height []int) int {
	if len(height) < 3 {
		return 0 // Not enough height elements to form a trapped area
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = common.Max(leftMax[i-1], height[i])
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = common.Max(rightMax[i+1], height[i])
	}

	water := 0
	for i := 1; i < len(height)-1; i++ {
		water += common.Min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

// O(n) time, O(1) space
func trapTwoPointer(height []int) int {
	if len(height) < 3 {
		return 0
	}
	l, r := 0, len(height)-1
	lMax, rMax, water := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			lMax = common.Max(lMax, height[l])
			water += lMax - height[l]
			l++
		} else {
			rMax = common.Max(rMax, height[r])
			water += rMax - height[r]
			r--
		}
	}
	return water
}

// Avg 2ns slower than original
func trapTwoPointerOpt(height []int) int {
	if len(height) < 3 {
		return 0
	}
	l, r := 0, len(height)-1
	lMax, rMax, water := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				water += lMax - height[l]
			}
			l++
		} else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				water += rMax - height[r]
			}
			r--
		}
	}
	return water
}

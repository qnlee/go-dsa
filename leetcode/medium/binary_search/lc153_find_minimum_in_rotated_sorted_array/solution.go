package lc153_find_minimum_in_rotated_sorted_array

import (
	"go-dsa/leetcode/helpers"
)

// Not optimized first attempt, not O(log n) time complexity
func findMinSoln1(nums []int) int {
	n := len(nums)
	if nums[0] <= nums[n-1] {
		return nums[0]
	}
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}

	return helpers.Min(nums[n-1], nums[n-2])
}

// Binary search, O(log n) time complexity
func findMinSoln2(nums []int) int {
	n := len(nums)
	if nums[0] <= nums[n-1] {
		return nums[0]
	}

	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}

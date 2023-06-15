package lc33_search_in_rotated_sorted_array

/*
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).

For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

Example 2:
	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1

Example 3:
	Input: nums = [1], target = 0
	Output: -1

Constraints:
	1 <= nums.length <= 5000
	-104 <= nums[i] <= 104
	All values of nums are unique.
	nums is an ascending array that is possibly rotated.
	-104 <= target <= 104
*/

// First attempt:
func searchTwoPass(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1
	// find pivot index (min element in nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	// use pivot to find target
	pivot := low
	low, high = 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		pMid := (mid + pivot) % n // adjust mid to account for pivot
		if nums[pMid] == target {
			return pMid
		} else if nums[pMid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// Optimized: pivot adjustment eliminated by directly checking the relative ordering of l, mid, r elements
func search1Pass(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target { // return target index if found
			return mid
		}
		if nums[l] <= nums[mid] { // ascending order
			if nums[l] <= target && target < nums[mid] { // l <= target < mid <= r
				r = mid - 1
			} else { // l <= mid <= target <= r
				l = mid + 1
			}
		} else { // not in ascending order
			if nums[mid] < target && target <= nums[r] { // l <= mid < target <= r
				l = mid + 1
			} else { // l <= target <= mid <= r
				r = mid - 1
			}
		}
	}

	return -1
}

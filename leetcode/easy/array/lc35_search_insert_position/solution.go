package lc35_search_insert_position

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return
the index where it would be if it were inserted in order. You must write an algorithm with O(log n) runtime complexity.

Example 1:
	Input: nums = [1,3,5,6], target = 5
	Output: 2
Example 2:
	Input: nums = [1,3,5,6], target = 2
	Output: 1
Example 3:
	Input: nums = [1,3,5,6], target = 7
	Output: 4

Constraints:
	1 <= nums.length <= 104
	-104 <= nums[i] <= 104
	nums contains distinct values sorted in ascending order.
	-104 <= target <= 104
*/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// At loop exit: l > r, and nums[r] < target < nums[l]
	// l is at the index of the smallest value in nums that is greater than target
	// r is at the index of the largest value in nums that is lower than target
	// so target would be inserted at index l or r+1
	return l
}

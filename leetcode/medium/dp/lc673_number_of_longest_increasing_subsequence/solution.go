package lc673_number_of_longest_increasing_subsequence

/*
Given an integer array nums, return the number of longest increasing subsequences. Notice that the sequence has to be
strictly increasing.

Example 1:
	Input: nums = [1,3,5,4,7]
	Output: 2
	Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
Example 2:
	Input: nums = [2,2,2,2,2]
	Output: 5
	Explanation: The length of the longest increasing subsequence is 1, and there are 5 increasing subsequences of
		length 1, so output 5.

Constraints:
	1 <= nums.length <= 2000
	-106 <= nums[i] <= 106
*/

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	lengths := make([]int, n) // lengths of LIS ending at each index.
	counts := make([]int, n)  // counts of LIS ending at each index.
	maxLIS, numLIS := 1, 0

	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
		// loop through the elements before the current index 'i'
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// if current element > element at j, we can extend the LIS
				// check if extending the LIS ending at j results in a longer LIS ending at i
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j] // if so, update the count with the count of LIS ending at j
				} else if lengths[j]+1 == lengths[i] {
					// if lengths are the same, add the count of LIS ending at j to the count of LIS ending at i
					counts[i] += counts[j]
				}
			}
		}

		if lengths[i] > maxLIS { // if LIS length up to index i is > previous max LIS length, update max/num to match
			maxLIS, numLIS = lengths[i], counts[i]
		} else if lengths[i] == maxLIS { // if LIS length up to i is same as previous max LIS length, add to total
			numLIS += counts[i]
		}
	}

	return numLIS
}

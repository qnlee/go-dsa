package lc673_number_of_longest_increasing_subsequence

// Function to find the number of Longest Increasing Subsequences (LIS) in a given slice of integers.
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	lengths := make([]int, n) // lengths of LIS ending at each index.
	counts := make([]int, n)  // counts of LIS ending at each index.
	maxLIS, numLIS := 1, 0

	// Loop through the 'nums' slice from the beginning.
	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
		// Loop through the elements before the current index 'i'.
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// If the current element is greater than the element at index 'j', it means we can extend the LIS.
				// Check if extending the LIS ending at index 'j' results in a longer LIS ending at index 'i'.
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j] // Update the count with the count of LIS ending at index 'j'.
				} else if lengths[j]+1 == lengths[i] {
					// If extending the LIS ending at index 'j' results in the same length of LIS ending at index 'i',
					// add the count of LIS ending at index 'j' to the count at index 'i'.
					counts[i] += counts[j]
				}
			}
		}
		// Update the overall maxLIS and numLIS.
		if lengths[i] > maxLIS {
			maxLIS, numLIS = lengths[i], counts[i]
		} else if lengths[i] == maxLIS {
			numLIS += counts[i]
		}
	}

	return numLIS
}

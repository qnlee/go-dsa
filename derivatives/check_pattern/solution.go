package check_pattern

/*
Given an input string and a pattern, check if characters in the input string follows the same order as determined by
characters present in the pattern. Assume there won’t be any duplicate characters in the pattern.

Examples:

	Input: string = "engineers rock", pattern = "er";
	Output: true
	All 'e' in the input string are before all 'r'.

	Input: string = "engineers rock", pattern = "egr";
	Output: false
	There are two 'e' after 'g' in the input string.

	Input: string = "engineers rock", pattern = "gsr";
	Output: false
	There are one 'r' before 's' in the input string.
*/

func checkPattern(str string, pat string) bool {
	label := make([]int, 256) // Assuming CHAR_SIZE is 256

	// Initialize all orders as -1
	for i := range label {
		label[i] = -1
	}

	// Assign an order to pattern characters according to their appearance in the pattern
	order := 1
	for i := 0; i < len(pat); i++ {
		// Assign the pattern characters order
		label[pat[i]] = order

		// Increment the order
		order++
	}

	// Now, check if string characters follow the given order
	lastOrder := -1
	for i := 0; i < len(str); i++ {
		if label[str[i]] != -1 {
			// If the order of this character is less than the order of the previous character, return false
			if label[str[i]] < lastOrder {
				return false
			}

			// Update lastOrder for the next iteration
			lastOrder = label[str[i]]
		}
	}

	// Return true if the string follows the given pattern
	return true
}

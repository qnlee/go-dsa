package lc3_longest_substring_without_repeating_characters

import "go-dsa/leetcode/helpers"

/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
	0 <= s.length <= 5 * 104
	s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	seen := make(map[byte]int)
	start, longest := 0, 1

	for i := 0; i < len(s); i++ {
		c := s[i]
		if prevIdx, ok := seen[c]; ok && prevIdx >= start {
			start = prevIdx + 1
		}
		seen[c] = i
		longest = helpers.Max(longest, i-start+1)
	}

	return longest
}

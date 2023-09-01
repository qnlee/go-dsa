package lc424_longest_repeating_character_replacement

import "go-dsa/leetcode/helpers"

/*
You are given a string s and an integer k. You can choose any character of the string and change it to any other
uppercase English character. You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achive this answer too.

Constraints:
1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
*/

func characterReplacement(s string, k int) int {
	charCounts := make([]int, 26)
	lo := 0
	longestCnt, maxLen := 0, 0
	for hi := 0; hi < len(s); hi++ {
		c := s[hi] - 'A'
		charCounts[c]++
		longestCnt = helpers.Max(longestCnt, charCounts[c])
		if longestCnt+k < hi-lo+1 {
			charCounts[s[lo]-'A']--
			lo++
		}
		maxLen = helpers.Max(maxLen, hi-lo+1)
	}
	return maxLen
}
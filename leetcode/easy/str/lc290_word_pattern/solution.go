package lc290_word_pattern

import "strings"

/*
Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:
	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true

Example 2:
	Input: pattern = "abba", s = "dog cat cat fish"
	Output: false

Example 3:
	Input: pattern = "aaaa", s = "dog cat cat dog"
	Output: false

Constraints:
	1 <= pattern.length <= 300
	pattern contains only lower-case English letters.
	1 <= s.length <= 3000
	s contains only lowercase English letters and spaces ' '.
	s does not contain any leading or trailing spaces.
	All the words in s are separated by a single space.
*/

func wordPattern(pattern string, s string) bool {
	// First split s to get slice of individual words
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	// To build map of word -> corresponding char
	charMatch := make(map[string]byte)
	// To keep track of which chars from pattern we have seen so far
	charSeen := make([]bool, 26)

	for i, word := range words {
		if char, exists := charMatch[word]; exists {
			if char != pattern[i] {
				return false
			}
			continue
		}

		// Ex. if pattern[i]='a', charIdx = 0
		charIdx := pattern[i] - 'a'
		// If we've seen the char but it doesn't exist in the word->char map
		if charSeen[charIdx] {
			return false
		}
		charSeen[charIdx] = true
		charMatch[word] = pattern[i]
	}

	return true
}

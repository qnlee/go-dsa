package lc125_valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	// Convert the string to lowercase
	s = strings.ToLower(s)
	// Remove non-alphanumeric characters
	filteredStr := make([]rune, 0, len(s))
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filteredStr = append(filteredStr, ch)
		}
	}

	// Check if the filtered string is a palindrome
	for left, right := 0, len(filteredStr)-1; left <= right; {
		if filteredStr[left] != filteredStr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

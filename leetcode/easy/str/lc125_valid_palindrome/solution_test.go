package lc125_valid_palindrome

import (
	"testing"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all
non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters
and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Ex1) Input: s = "A man, a plan, a canal: Panama"
	 Output: true
	 Explanation: "amanaplanacanalpanama" is a palindrome.

Ex2) Input: s = "race a car"
	 Output: false
	 Explanation: "raceacar" is not a palindrome.

Ex3) Input: s = " "
 	 Output: true
	 Explanation: s is an empty string "" after removing non-alphanumeric characters.
	 Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
	1 <= s.length <= 2 * 105
	s consists only of printable ASCII characters.
*/

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in       string
		expected bool
	}{
		{in: "A man, a plan, a canal: Panama", expected: true},
		{in: "race a car", expected: false},
		{in: " ", expected: true},
		{in: "", expected: true},
		{in: "A plan, a man, a canal: Panama", expected: false},
		{in: "aaabaaa", expected: true},
		{in: "aabaaa", expected: false},
	}

	for _, c := range cases {
		actual := isPalindrome(c.in)
		if actual != c.expected {
			t.Errorf("isPalindrome(%s): Expected [%t], but got [%t]", c.in, c.expected, actual)
		}
	}
}

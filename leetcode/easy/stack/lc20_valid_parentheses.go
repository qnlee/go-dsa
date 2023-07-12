package stack

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.

Example 1:
	Input: s = "()"
	Output: true
Example 2:
	Input: s = "()[]{}"
	Output: true
Example 3:
	Input: s = "(]"
	Output: false

Constraints:
	1 <= s.length <= 104
	s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		switch c {
		case '{':
			stack = append(stack, '}')
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if c != pop {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

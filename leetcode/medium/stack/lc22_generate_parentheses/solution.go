package lc22_generate_parentheses

func generateParentheses(n int) []string {
	var res []string
	helper(n, 0, 0, "", &res)
	return res
}

func helper(n, left, right int, curr string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, curr)
	}
	if left < n {
		helper(n, left+1, right, curr+"(", res)
	}
	if right < n {
		helper(n, left, right+1, curr+")", res)
	}
}

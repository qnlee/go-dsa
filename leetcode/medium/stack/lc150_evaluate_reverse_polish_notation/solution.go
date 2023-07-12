package lc150_evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		n := len(stack)
		switch tokens[i] {
		case "+":
			stack[n-2] = stack[n-2] + stack[n-1]
			stack = stack[:n-1]
		case "-":
			stack[n-2] = stack[n-2] - stack[n-1]
			stack = stack[:n-1]
		case "*":
			stack[n-2] = stack[n-2] * stack[n-1]
			stack = stack[:n-1]
		case "/":
			stack[n-2] = stack[n-2] / stack[n-1]
			stack = stack[:n-1]
		default: // int value
			v, _ := strconv.Atoi(tokens[i])
			stack = append(stack, v)
		}
	}
	return stack[0]
}

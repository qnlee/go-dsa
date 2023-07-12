package lc739_daily_temperatures

/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i]
is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for
which this is possible, keep answer[i] == 0 instead.

Example 1:
	Input: temperatures = [73,74,75,71,69,72,76,73]
	Output: [1,1,4,2,1,1,0,0]
Example 2:
	Input: temperatures = [30,40,50,60]
	Output: [1,1,1,0]
Example 3:
	Input: temperatures = [30,60,90]
	Output: [1,1,0]

Constraints:
	1 <= temperatures.length <= 105
	30 <= temperatures[i] <= 100
*/

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			ans[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

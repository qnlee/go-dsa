package lc70_climbing_stairs

/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:
1 <= n <= 45
*/

// Thought process:
// from step 0 to step 1: only 1
// from step 1 to step 2: only 1 (+ ways to get to step 1 from step 0)
// from step 2 to step 3: only 1 (+ ways to get to step 2 from step 1)
// from step 3 to step 4: only 1 (+ ways to get to step 3 from step 2) // n=1: 1 way
// n=2: 1+1 / 2 == 2 ways
// n=3: 1+1+1 / 1+2 / 2+1 == 3 ways
// n=4: 2+2 / 2+1+1 / 1+2+1 / 1+1+2 / 1+1+1+1 == 5 ways == climbStairs(4-1) + climbStairs(4-2)

func climbStairs(n int, soln string) int {
	switch soln {
	case "recursive":
		return recursive(n)
	case "recursiveMemo":
		return recursiveMemo(n)
	case "bottomUp":
		return bottomUp(n)
	case "bottomUpOptimized":
		return bottomUpOptimized(n)
	default:
		return bottomUpOptimized(n)
	}
}

// Top-down approach (recursion)
// Time complexity: O(2^n)
// Space complexity: O(n)
func recursive(n int) int {
	if n < 3 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

// Optimization: Reduce need for repeated executions by storing prev results in cache ("memoization")
var cache []int

// Time complexity: O(n)
// Space complexity: O(n+1) = O(n)
func recursiveMemo(n int) int {
	cache = make([]int, n+1)
	return recursiveMemoHelper(n)
}

func recursiveMemoHelper(n int) int {
	if n < 3 {
		return n
	}
	if cache[n] != 0 {
		return cache[n]
	}

	res := recursiveMemoHelper(n-1) + recursiveMemoHelper(n-2)
	cache[n] = res
	return res
}

// Bottom-up approach: Since we know the steps for n=1,2,3 we can use those to iteratively build up a solution
// Time complexity: O(n-3) = O(n)
// Space complexity: O(n)
func bottomUp(n int) int {
	if n < 3 {
		return n
	}

	ways := make([]int, n+1)
	ways[1], ways[2], ways[3] = 1, 2, 3

	for i := 4; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}

// Optimization: We don't even need all of the previous results-- we can just keep track of the previous 2 results
// Time complexity: O(n)
// Space complexity: O(1)
func bottomUpOptimized(n int) int {
	if n < 3 {
		return n
	}
	// eg) say we are on step 3 (3 ways to get to step 3), and previous step was step 2 (2 ways)
	curr, prev := 3, 2
	for i := 4; i <= n; i++ {
		// eg) set prev <- steps to get to step 3,
		//     curr <- steps to get to step 4 == ways to get to step 3 + ways to get to step 2
		prev, curr = curr, curr+prev
	}

	return curr
}

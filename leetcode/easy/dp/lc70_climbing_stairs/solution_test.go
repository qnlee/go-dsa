package lc70_climbing_stairs

import (
	"testing"
)

func TestClimbStairs_AllSolutions(t *testing.T) {
	cases := []struct {
		n        int
		soln     string
		expected int
	}{
		{n: 4, soln: "recursive", expected: 5},
		{n: 4, soln: "recursiveMemo", expected: 5},
		{n: 4, soln: "bottomUp", expected: 5},
		{n: 4, soln: "bottomUpOptimized", expected: 5},
	}

	for _, c := range cases {
		actual := climbStairs(c.n, c.soln)
		if actual != c.expected {
			t.Errorf(
				"ClimbStairs(%d) - %s solution: Expected [%d], but got [%d]",
				c.n, c.soln, c.expected, actual,
			)
		}
	}
}

func BenchmarkClimbStairs_Recursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		climbStairs(4, "recursive")
	}
}
func BenchmarkClimbStairs_RecursiveMemo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		climbStairs(4, "recursiveMemo")
	}
}
func BenchmarkClimbStairs_BottomUp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		climbStairs(4, "bottomUp")
	}
}
func BenchmarkClimbStairs_BottomUpOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		climbStairs(4, "bottomUpOptimized")
	}
}

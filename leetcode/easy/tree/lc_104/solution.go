package lc_104

import (
	"go-dsa/leetcode/easy/tree"
	"go-dsa/leetcode/helpers"
)

func maxDepthRecursive(root *tree.Node) int {
	if root == nil {
		return 0
	}
	return 1 + helpers.Max(maxDepthRecursive(root.Left), maxDepthRecursive(root.Right))
}

func maxDepthIterative(root *tree.Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	level := []*tree.Node{root}
	for len(level) != 0 {
		depth += 1
		var nextLevel []*tree.Node
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
	}
	return depth
}

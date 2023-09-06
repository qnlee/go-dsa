package lc104_maximum_depth_of_binary_tree

import (
	"go-dsa/leetcode/common"
)

func maxDepthRecursive(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + common.Max(maxDepthRecursive(root.Left), maxDepthRecursive(root.Right))
}

func maxDepthIterative(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	level := []*common.TreeNode{root}
	for len(level) != 0 {
		depth += 1
		var nextLevel []*common.TreeNode
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

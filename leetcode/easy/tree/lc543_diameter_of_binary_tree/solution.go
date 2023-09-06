package lc543_diameter_of_binary_tree

import (
	"go-dsa/leetcode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *common.TreeNode) int {
	var res int
	var maxDepth func(node *common.TreeNode) int
	maxDepth = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := maxDepth(node.Left), maxDepth(node.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}

	maxDepth(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

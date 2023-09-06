package lc226_invert_binary_tree

import (
	"go-dsa/leetcode/common"
)

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

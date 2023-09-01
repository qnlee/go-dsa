package lc226_invert_binary_tree

import "go-dsa/leetcode/easy/tree"

func invertTree(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

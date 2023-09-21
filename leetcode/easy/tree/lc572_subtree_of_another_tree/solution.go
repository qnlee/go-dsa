package lc572_subtree_of_another_tree

import "go-dsa/leetcode/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *common.TreeNode, subRoot *common.TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == subRoot.Val && sameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}

package common

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Insert(val int) (*TreeNode, error) {
	if tn == nil {
		return nil, errors.New("cannot insert value into nil tree")
	}
	switch {
	case val == tn.Val:
		return nil, nil
	case val < tn.Val:
		if tn.Left == nil {
			tn.Left = &TreeNode{Val: val}
			return tn.Left, nil
		}
		return tn.Left.Insert(val)
	default:
		if tn.Right == nil {
			tn.Right = &TreeNode{Val: val}
			return tn.Right, nil
		}
		return tn.Right.Insert(val)
	}
}

func (tn *TreeNode) ToString() string {
	if tn == nil {
		return "null"
	}
	leftStr := tn.Left.ToString()
	rightStr := tn.Right.ToString()
	nodeStr := fmt.Sprintf("{Val: %d, Left: %s, Right: %s}", tn.Val, leftStr, rightStr)

	return nodeStr
}

func (tn *TreeNode) SameTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root == nil {
		return false
	}
	return tn.Val == root.Val && tn.Left.SameTree(root.Left) && tn.Right.SameTree(root.Right)
}

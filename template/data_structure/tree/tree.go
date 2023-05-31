package tree

import (
	"errors"
)

type Node struct {
	Val      int
	Children []*Node
}

type BNode struct {
	Val   int
	Left  *BNode
	Right *BNode
}

type BSTree struct {
	Root *BNode
}

func NewNode(val int) *Node {
	return &Node{
		Val:      val,
		Children: []*Node{},
	}
}

func (n *Node) Insert(val int) {
	n.Children = append(n.Children, NewNode(val))
}

func (n *Node) RemoveChild(target *Node) {
	for i, c := range n.Children {
		if c == target {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			break
		}
	}
}

func NewBNode(val int) *BNode {
	return &BNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (tn *BNode) Insert(val int) (*BNode, error) {
	if tn == nil {
		return nil, errors.New("cannot insert value into nil tree")
	}

	switch {
	case val == tn.Val:
		return nil, nil
	case val < tn.Val:
		if tn.Left == nil {
			tn.Left = NewBNode(val)
			return tn.Left, nil
		}
		return tn.Left.Insert(val)
	default:
		if tn.Right == nil {
			tn.Right = NewBNode(val)
			return tn.Right, nil
		}
		return tn.Right.Insert(val)
	}
}

func (tn *BNode) Find(target int) *BNode {
	if tn == nil {
		return nil
	}

	switch {
	case tn.Val == target:
		return tn
	case tn.Val < target:
		return tn.Left.Find(target)
	default:
		return tn.Right.Find(target)
	}
}

func (tn *BNode) PreOrderWalk(ch chan int) {
	if tn == nil {
		return
	}
	ch <- tn.Val
	if tn.Left != nil {
		tn.Left.PreOrderWalk(ch)
	}
	if tn.Right != nil {
		tn.Right.PreOrderWalk(ch)
	}
}

func (tn *BNode) InOrderWalk(ch chan int) {
	if tn == nil {
		return
	}
	if tn.Left != nil {
		tn.Left.InOrderWalk(ch)
	}
	ch <- tn.Val
	if tn.Right != nil {
		tn.Right.InOrderWalk(ch)
	}
}

func (tn *BNode) PostOrderWalk(ch chan int) {
	if tn == nil {
		return
	}
	if tn.Left != nil {
		tn.Left.PostOrderWalk(ch)
	}
	if tn.Right != nil {
		tn.Right.PostOrderWalk(ch)
	}
	ch <- tn.Val
}

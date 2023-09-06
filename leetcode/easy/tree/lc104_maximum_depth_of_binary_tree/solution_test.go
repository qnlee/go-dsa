package lc104_maximum_depth_of_binary_tree

import (
	"go-dsa/leetcode/common"
	"testing"
)

func TestMaxDepthRecursive(t *testing.T) {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  &common.TreeNode{Val: 4},
			Right: &common.TreeNode{Val: 5},
		},
		Right: &common.TreeNode{Val: 3},
	}

	expectedDepth := 3
	actualDepth := maxDepthRecursive(root)
	if actualDepth != expectedDepth {
		t.Errorf("Expected depth: %d, but got: %d", expectedDepth, actualDepth)
	}
}

func TestMaxDepthIterative(t *testing.T) {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  &common.TreeNode{Val: 4},
			Right: &common.TreeNode{Val: 5},
		},
		Right: &common.TreeNode{Val: 3},
	}

	expectedDepth := 3
	actualDepth := maxDepthIterative(root)
	if actualDepth != expectedDepth {
		t.Errorf("Expected depth: %d, but got: %d", expectedDepth, actualDepth)
	}
}

func BenchmarkMaxDepthRecursive(b *testing.B) {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 4,
			},
			Right: &common.TreeNode{
				Val: 5,
			},
		},
		Right: &common.TreeNode{
			Val: 3,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepthRecursive(root)
	}
}

func BenchmarkMaxDepthIterative(b *testing.B) {
	root := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 4,
			},
			Right: &common.TreeNode{
				Val: 5,
			},
		},
		Right: &common.TreeNode{
			Val: 3,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepthIterative(root)
	}
}

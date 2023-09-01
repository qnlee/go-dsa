package lc_104

import (
	"go-dsa/leetcode/easy/tree"
	"testing"
)

func TestMaxDepthRecursive(t *testing.T) {
	root := &tree.Node{
		Val: 1,
		Left: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val: 4,
			},
			Right: &tree.Node{
				Val: 5,
			},
		},
		Right: &tree.Node{
			Val: 3,
		},
	}

	expectedDepth := 3
	actualDepth := maxDepthRecursive(root)
	if actualDepth != expectedDepth {
		t.Errorf("Expected depth: %d, but got: %d", expectedDepth, actualDepth)
	}
}

func TestMaxDepthIterative(t *testing.T) {
	root := &tree.Node{
		Val: 1,
		Left: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val: 4,
			},
			Right: &tree.Node{
				Val: 5,
			},
		},
		Right: &tree.Node{
			Val: 3,
		},
	}

	expectedDepth := 3
	actualDepth := maxDepthIterative(root)
	if actualDepth != expectedDepth {
		t.Errorf("Expected depth: %d, but got: %d", expectedDepth, actualDepth)
	}
}

func BenchmarkMaxDepthRecursive(b *testing.B) {
	root := &tree.Node{
		Val: 1,
		Left: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val: 4,
			},
			Right: &tree.Node{
				Val: 5,
			},
		},
		Right: &tree.Node{
			Val: 3,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepthRecursive(root)
	}
}

func BenchmarkMaxDepthIterative(b *testing.B) {
	root := &tree.Node{
		Val: 1,
		Left: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val: 4,
			},
			Right: &tree.Node{
				Val: 5,
			},
		},
		Right: &tree.Node{
			Val: 3,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepthIterative(root)
	}
}

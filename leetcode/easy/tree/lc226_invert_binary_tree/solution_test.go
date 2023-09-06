package lc226_invert_binary_tree

import (
	"fmt"
	"go-dsa/leetcode/easy/tree"
	"testing"
)

func TestInvertTree(t *testing.T) {
	testTrees, err := tree.LoadExampleTrees("../test_trees.json")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := tree.LoadExampleTrees("../test_trees_inverted.json")
	if err != nil {
		t.Fatal(err)
	}

	for i, testTree := range testTrees {
		t.Run(fmt.Sprintf("test tree %d", i+1), func(t *testing.T) {
			actual := invertTree(&testTree.TreeNode)
			expected := &expected[i].TreeNode
			if !actual.SameTree(expected) {
				t.Errorf("Expected inverted tree: %s, but got: %s", actual.ToString(), expected.ToString())
			}
		})
	}
}

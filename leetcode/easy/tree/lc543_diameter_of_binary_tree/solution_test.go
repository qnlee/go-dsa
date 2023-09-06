package lc543_diameter_of_binary_tree

import (
	"fmt"
	"go-dsa/leetcode/easy/tree"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	testTrees, err := tree.LoadExampleTrees("../test_trees.json")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		tree     tree.ExampleTree
		expected int
	}{
		{
			//        1
			//       / \
			//      2   3
			//     / \
			//    4   5
			tree:     testTrees[0],
			expected: 3,
		}, {
			//        15
			//        / \
			//       20  7
			tree:     testTrees[1],
			expected: 2,
		}, {
			//        1
			//         \
			//          2
			//           \
			//            3
			tree:     testTrees[2],
			expected: 2,
		}, {
			//        1
			//       /
			//      2
			tree:     testTrees[3],
			expected: 1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test tree %d", i+1), func(t *testing.T) {
			result := diameterOfBinaryTree(&testCase.tree.TreeNode)
			if result != testCase.expected {
				t.Errorf("Expected diameter: %d, but got: %d", testCase.expected, result)
			}
		})
	}
}

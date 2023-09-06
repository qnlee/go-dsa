package tree

import (
	"encoding/json"
	"fmt"
	"go-dsa/leetcode/common"
	"os"
)

type ExampleTree struct {
	common.TreeNode `json:"tree"`
}

func LoadExampleTrees(filename string) ([]ExampleTree, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return []ExampleTree{}, fmt.Errorf("error reading %s: %v", filename, err)
	}

	var trees []ExampleTree
	if err := json.Unmarshal(fileData, &trees); err != nil {
		return []ExampleTree{}, fmt.Errorf("error parsing JSON in file %s: %v", filename, err)
	}

	return trees, nil
}

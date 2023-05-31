package tree

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	root := NewBNode(4)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)
	root.Insert(5)
	root.Insert(7)

	preOrderCh := make(chan int)
	go func() {
		defer close(preOrderCh)
		root.PreOrderWalk(preOrderCh)
	}()
	// 4  2  1  3  6  5  7
	fmt.Println("Pre-order traversal:")
	for val := range preOrderCh {
		fmt.Printf(" %d ", val)
	}

	inOrderCh := make(chan int)
	go func() {
		defer close(inOrderCh)
		root.InOrderWalk(inOrderCh)
	}()
	//  1  2  3  4  5  6  7
	fmt.Println("\nIn-order traversal:")
	for val := range inOrderCh {
		fmt.Printf(" %d ", val)
	}

	postOrderCh := make(chan int)
	go func() {
		defer close(postOrderCh)
		root.PostOrderWalk(postOrderCh)
	}()
	i := 0
	expected := []int{1, 3, 2, 5, 7, 6, 4}
	fmt.Println("\nPost-order traversal:")
	for val := range postOrderCh {
		fmt.Printf(" %d ", val)
		if expected[i] != val {
			t.Errorf("PostOrder traversal: expected [%d], actual [%d]", expected[i], val)
		}
		i += 1
	}
}

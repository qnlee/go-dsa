package search_sort_test

import (
	"fmt"
	"go-dsa/templates/search_sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.BubbleSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !isEqual(arr, expected) {
		t.Errorf("BubbleSort: Expected %v, but got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.SelectionSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !isEqual(arr, expected) {
		t.Errorf("SelectionSort: Expected %v, but got %v", expected, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.InsertionSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !isEqual(arr, expected) {
		t.Errorf("InsertionSort: Expected %v, but got %v", expected, arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.MergeSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !isEqual(arr, expected) {
		t.Errorf("MergeSort: Expected %v, but got %v", expected, arr)
	}
}

func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func ExampleBubbleSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.BubbleSort(arr)
	fmt.Println(arr)
	// Output: [11 12 22 25 34 64 90]
}

func ExampleSelectionSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.SelectionSort(arr)
	fmt.Println(arr)
	// Output: [11 12 22 25 34 64 90]
}

func ExampleInsertionSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.InsertionSort(arr)
	fmt.Println(arr)
	// Output: [11 12 22 25 34 64 90]
}

func ExampleMergeSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	search_sort.MergeSort(arr)
	fmt.Println(arr)
	// Output: [11 12 22 25 34 64 90]
}

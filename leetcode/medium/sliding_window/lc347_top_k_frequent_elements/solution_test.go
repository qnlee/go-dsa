package lc347_top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"
)

func TestSolution_TopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3, 3, 3}, 3, []int{1, 2, 3}},
	}

	for _, test := range tests {
		got := topKFrequent(test.nums, test.k)
		sort.Ints(got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("TopKFrequent(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
		}
	}
}

func TestSolution_TopKFrequentMinHeap(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3, 3, 3}, 3, []int{1, 2, 3}},
	}

	for _, test := range tests {
		got := topKFrequentMinHeap(test.nums, test.k)
		sort.Ints(got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("TopKFrequentMinHeap(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
		}
	}
}

func BenchmarkSolution_TopKFrequent(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	for i := 0; i < b.N; i++ {
		topKFrequent(nums, k)
	}
}

func BenchmarkSolution_TopKFrequentMinHeap(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	for i := 0; i < b.N; i++ {
		topKFrequentMinHeap(nums, k)
	}
}

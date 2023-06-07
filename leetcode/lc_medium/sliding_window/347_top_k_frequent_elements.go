package sliding_window

import (
	"container/heap"
	"go-dsa/template/heap_queue"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

Example 2:
	Input: nums = [1], k = 1
	Output: [1]

Constraints:
	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
	k is in the range [1, the number of unique elements in the array].
	It is guaranteed that the answer is unique.
*/

func TopKFrequent(nums []int, k int) []int {
	// count frequencies of unique elements in nums
	freqs := make(map[int]int)
	for _, n := range nums {
		freqs[n] += 1
	}
	// use buckets (rather than just []int, etc.) to store elements by frequency, since nums could contain multiple
	// elements that occur at the same frequencies
	buckets := make([][]int, len(nums)+1)
	for n, f := range freqs {
		buckets[f] = append(buckets[f], n)
	}
	// add elements from buckets starting at highest frequency bucket until we have k most frequent elements
	var res []int
	for i := len(buckets) - 1; len(res) < k && i >= 0; i-- {
		if buckets[i] == nil {
			continue
		}
		res = append(res, buckets[i]...)
	}
	return res
}

func TopKFrequentMinHeap(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, n := range nums {
		freqs[n]++
	}

	h := &heap_queue.MinHeap{}
	heap.Init(h)

	for n, f := range freqs {
		heap.Push(h, &heap_queue.Element{Value: n, Priority: f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(heap_queue.Element).Value
	}

	return res
}

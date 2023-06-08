package heap_queue

import "container/heap"

// Based on example of PriorityQueue taken from https://pkg.go.dev/container/heap

type QueueElement struct {
	Value    int
	Priority int
	Index    int
}

type PriorityQueue []*QueueElement

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	elem := x.(*QueueElement)
	elem.Index = n
	*pq = append(*pq, elem)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	elem := old[n-1]
	old[n-1] = nil  // avoid memory leak
	elem.Index = -1 // for safety
	*pq = old[0 : n-1]
	return elem
}

func (pq *PriorityQueue) update(elem *QueueElement, value int, priority int) {
	elem.Value = value
	elem.Priority = priority
	heap.Fix(pq, elem.Index)
}

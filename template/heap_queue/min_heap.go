package heap_queue

// Based on example of IntHeap taken from https://pkg.go.dev/container/heap

type HeapElement struct {
	Value     int
	Frequency int
}

type MinHeap []HeapElement

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Frequency < mh[j].Frequency
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(HeapElement))
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

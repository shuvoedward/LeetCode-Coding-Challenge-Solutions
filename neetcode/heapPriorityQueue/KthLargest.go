package heappriorityqueue

import "container/heap"

type minheap []int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

type kthLargest struct {
	k    int
	heap minheap
}

func constructor(k int, nums []int) kthLargest {
	kth := kthLargest{k, minheap{}}

	for _, num := range nums {
		heap.Push(&kth.heap, num)
	}

	return kth
}

func (this *kthLargest) Add(val int) int {
	heap.Push(&this.heap, val)

	if this.heap.Len() > this.k {
		heap.Pop(&this.heap)
	}

	return this.heap[0]
}

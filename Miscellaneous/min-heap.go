package Miscellaneous

import "container/heap"

//
// time complexity: O(klogn + n)
// space complexity: O(n)
//
func findKthSmallestNumberUsingMinHeap(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	var minHeap MinHeap
	heap.Init(&minHeap)

	for i := 0; i < len(nums); i++ {
		heap.Push(&minHeap, nums[i])
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(&minHeap)
	}

	return minHeap[0]
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

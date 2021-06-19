package Miscellaneous

import "container/heap"

//
// time complexity: O(klogk + (n-k)log(k))
// space complexity: O(k)
//
func findKthSmallestNumberUsingMaxHeap(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	var maxheap MaxHeap
	heap.Init(&maxheap)

	for i := 0; i < k; i++ {
		heap.Push(&maxheap, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < maxheap[0] {
			heap.Pop(&maxheap)
			heap.Push(&maxheap, nums[i])
		}
	}

	return maxheap[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

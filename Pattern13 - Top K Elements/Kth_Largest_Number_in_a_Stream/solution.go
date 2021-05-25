package kthlargestnumberinastream

import "container/heap"

/*
Design a class to efficiently find the Kth largest element in a stream of numbers.

The class should have the following two things:
1. The constructor of the class should accept an integer array containing initial numbers from the stream and an integer ‘K’.
2. The class should expose a function add(int num) which will store the given number and return the Kth largest number.

Input: [3, 1, 5, 12, 2, 11], K = 4
1. Calling add(6) should return '5'.
2. Calling add(13) should return '6'.
2. Calling add(4) should still return '6'.

ref: https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
*/

type KthLargest struct {
	H IntHeap
	K int
}

func Constructor(k int, nums []int) KthLargest {
	model := KthLargest{}
	model.K = k
	heap.Init(&model.H)

	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(&model.H, nums[i])
		} else if nums[i] > model.H[0] {
			heap.Pop(&model.H)
			heap.Push(&model.H, nums[i])
		}
	}

	return model
}

func (this *KthLargest) Add(val int) int {
	if this.H.Len() < this.K {
		heap.Push(&this.H, val)
		if this.H.Len() == this.K {
			return this.H[0]
		}
	}

	if val > this.H[0] {
		heap.Pop(&this.H)
		heap.Push(&this.H, val)
	}

	return this.H[0]
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

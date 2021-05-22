package kthsmallestnumber

import "container/heap"

/*
Given an unsorted array of numbers, find Kth smallest number in it.
Please note that it is the Kth smallest number in the sorted order, not the Kth distinct element.
Note: For a detailed discussion about different approaches to solve this problem, take a look at Kth Smallest Number.
Kth Smallest Number: https://www.educative.io/courses/grokking-the-coding-interview/myJK6Wvj00R

Input: [1, 5, 12, 2, 11, 5], K = 3
Output: 5
Explanation: The 3rd smallest number is '5', as the first two smaller numbers are [1, 2].

Input: [1, 5, 12, 2, 11, 5], K = 4
Output: 5
Explanation: The 4th smallest number is '5', as the first three small numbers are [1, 2, 5].

Input: [5, 12, 11, -1, 12], K = 3
Output: 11
Explanation: The 3rd smallest number is '11', as the first two small numbers are [5, -1].

ref: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthSmallest(nums []int, k int) int {
	var (
		h IntHeap
		i = 0
	)
	for ; i < k; i++ {
		h = append(h, nums[i])
	}
	heap.Init(&h)
	for ; i < len(nums); i++ {
		if nums[i] < h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}

	return h[0]
}

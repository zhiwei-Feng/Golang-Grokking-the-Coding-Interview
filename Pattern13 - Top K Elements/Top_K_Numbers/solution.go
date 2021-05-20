package topknumbers

import "container/heap"

/*
Given an unsorted array of numbers, find the ‘K’ largest numbers in it.

Note: For a detailed discussion about different approaches to solve this problem, take a look at Kth Smallest Number.

Input: [3, 1, 5, 12, 2, 11], K = 3
Output: [5, 12, 11]

Input: [5, 12, 11, -1, 12], K = 3
Output: [12, 11, 12]

ref: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func getLargestNumbers(arr []int, k int) []int {
	var (
		intHeap IntHeap
		i       = 0
	)
	for ; i < k; i++ {
		intHeap = append(intHeap, arr[i])
	}
	heap.Init(&intHeap)
	for ; i < len(arr); i++ {
		if arr[i] > intHeap[0] {
			heap.Push(&intHeap, arr[i])
			heap.Pop(&intHeap)
		}
	}

	return intHeap
}

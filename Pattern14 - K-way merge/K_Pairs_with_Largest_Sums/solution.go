package kpairswithlargestsums

import "container/heap"

/*
* K Pairs with Largest Sums
Given two sorted arrays in descending order, find ‘K’ pairs with the largest sum where each pair consists of numbers from both the arrays.

Input: L1=[9, 8, 2], L2=[6, 3, 1], K=3
Output: [9, 3], [9, 6], [8, 6]
Explanation: These 3 pairs have the largest sum. No other pair has a sum larger than any of these.

Input: L1=[5, 2, 1], L2=[2, -1], K=3
Output: [5, 2], [5, -1], [2, 2]

ref: https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/
*/

func kLargestPairs(nums1, nums2 []int, k int) [][]int {
	var (
		mheap MinHeap
	)
	heap.Init(&mheap)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if mheap.Len() < k {
				heap.Push(&mheap, []int{v1, v2})
			} else if v1+v2 > mheap[0][0]+mheap[0][1] {
				heap.Pop(&mheap)
				heap.Push(&mheap, []int{v1, v2})
			} else {
				break
			}
		}
	}

	return mheap
}

type MinHeap [][]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return (h[i][0] + h[i][1]) < (h[j][0] + h[j][1]) }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

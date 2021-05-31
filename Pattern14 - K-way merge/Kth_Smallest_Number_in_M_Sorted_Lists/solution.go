package kthsmallestnumberinmsortedlists

import "container/heap"

/*
* Kth Smallest Number in M Sorted Lists
Given ‘M’ sorted arrays, find the K’th smallest number among all the arrays.

Input: L1=[2, 6, 8], L2=[3, 6, 7], L3=[1, 3, 4], K=5
Output: 4
Explanation: The 5th smallest number among all the arrays is 4, this can be verified from
the merged list of all the arrays: [1, 2, 3, 3, 4, 6, 6, 7, 8]

Input: L1=[5, 8, 9], L2=[1, 7], K=3
Output: 7
Explanation: The 3rd smallest number among all the arrays is 7.

ref: https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
*/

func kthSmallest(matrix [][]int, k int) int {
	var (
		count = 0
		h     MinHeap
	)
	heap.Init(&h)
	for i, arr := range matrix {
		if len(arr) > 0 {
			heap.Push(&h, Element{arr[0], i, 0})
		}
	}

	for h.Len() > 0 {
		curSmalEle := heap.Pop(&h).(Element)
		count++
		if count == k {
			return curSmalEle.Val
		}
		belongArr := matrix[curSmalEle.ArrInd]
		nextInd := curSmalEle.EleInd + 1
		if nextInd < len(belongArr) {
			heap.Push(&h, Element{belongArr[nextInd], curSmalEle.ArrInd, nextInd})
		}
	}

	return -1
}

type Element struct {
	Val    int
	ArrInd int
	EleInd int
}

type MinHeap []Element

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Element)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

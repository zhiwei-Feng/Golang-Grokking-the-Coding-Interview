package kthsmallestnumberinasortedmatrix

import "container/heap"

/*
* Kth Smallest Number in a Sorted Matrix
Given an N * NN∗N matrix where each row and column is sorted in ascending order, find the Kth smallest element in the matrix.

Input: Matrix=[
    [2, 6, 8],
    [3, 7, 10],
    [5, 8, 11]
  ],
  K=5
Output: 7
Explanation: The 5th smallest number in the matrix is 7.

ref: https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
*/

func kthSmallest(matrix [][]int, k int) int {
	var (
		count = 0
		h     MinHeap
	)
	heap.Init(&h)
	n := len(matrix)
	// 这里注意不一定需要将所有行考虑进去，因为k可以小于n
	for i := 0; i < n && i < k; i++ {
		if arr := matrix[i]; len(arr) > 0 {
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

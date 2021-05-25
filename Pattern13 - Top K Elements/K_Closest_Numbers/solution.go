package kclosestnumbers

import (
	"container/heap"
	"sort"
)

/*
Given a sorted number array and two integers ‘K’ and ‘X’, find ‘K’ closest numbers to ‘X’ in the array.
Return the numbers in the sorted order. ‘X’ is not necessarily present in the array.

Input: [5, 6, 7, 8, 9], K = 3, X = 7
Output: [6, 7, 8]

Input: [2, 4, 5, 6, 9], K = 3, X = 6
Output: [4, 5, 6]

Input: [2, 4, 5, 6, 9], K = 3, X = 10
Output: [5, 6, 9]

ref: https://leetcode-cn.com/problems/find-k-closest-elements/
*/

func findClosestElements(arr []int, k int, x int) []int {
	index := binarySearch(arr, x)
	low, high := index-k, index+k
	if low < 0 {
		low = 0
	}
	if high > len(arr)-1 {
		high = len(arr) - 1
	}
	var h IntHeap
	heap.Init(&h)

	for i := low; i <= high; i++ {
		heap.Push(&h, Integer{Value: arr[i], Distance: abs(arr[i] - x)})
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(Integer).Value)
	}

	sort.Ints(res)
	return res
}

func binarySearch(arr []int, x int) int {
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == x {
			return mid
		}

		if x < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if start > 0 {
		return start - 1
	} else {
		return start
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Integer struct {
	Value    int
	Distance int
}
type IntHeap []Integer

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	if h[i].Distance != h[j].Distance {
		return h[i].Distance < h[j].Distance
	} else {
		return h[i].Value < h[j].Value
	}
}
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(Integer)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

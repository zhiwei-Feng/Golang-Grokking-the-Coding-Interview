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
	var (
		intMap IntMap
	)

	heap.Init(&intMap)
	for i := 0; i < len(arr); i++ {
		if intMap.Len() < k {
			heap.Push(&intMap, Integer{arr[i], distance(arr[i], x)})
			continue
		}

		curDis := distance(x, arr[i])
		if curDis < intMap[0].Distance || (curDis == intMap[0].Distance && arr[i] < intMap[0].Value) {
			heap.Pop(&intMap)
			heap.Push(&intMap, Integer{arr[i], curDis})
		}
	}

	res := make([]int, 0, len(intMap))
	for _, v := range intMap {
		res = append(res, v.Value)
	}
	sort.Ints(res)
	return res
}

func distance(x, y int) int {
	res := x - y
	if res < 0 {
		res = -res
	}
	return res
}

type Integer struct {
	Value    int
	Distance int
}
type IntMap []Integer

func (h IntMap) Len() int            { return len(h) }
func (h IntMap) Less(i, j int) bool  { return h[i].Distance > h[j].Distance }
func (h IntMap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMap) Push(x interface{}) { *h = append(*h, x.(Integer)) }
func (h *IntMap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

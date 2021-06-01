package smallestnumberrange

import (
	"container/heap"
	"math"
)

/*
* Smallest Number Range
Given ‘M’ sorted arrays, find the smallest range that includes at least one number from each of the ‘M’ lists.

Input: L1=[1, 5, 8], L2=[4, 12], L3=[7, 8, 10]
Output: [4, 7]
Explanation: The range [4, 7] includes 5 from L1, 4 from L2 and 7 from L3.

Input: L1=[1, 9], L2=[4, 12], L3=[7, 10, 16]
Output: [9, 12]
Explanation: The range [9, 12] includes 9 from L1, 12 from L2 and 10 from L3.

ref: https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/
*/

func smallestRange(nums [][]int) []int {
	//
	// 关键点
	// 1. 随时跟踪当前heap中最大的元素
	// 2. 保持heap的大小为len(nums), 在这个情况下说明当前的range可以满足条件
	//
	var (
		h             MinHeap
		currentMaxNum = math.MinInt32
		rangeStart    = 0
		rangeEnd      = math.MaxInt32
	)
	heap.Init(&h)
	for i := 0; i < len(nums); i++ {
		if len(nums[i]) > 0 {
			heap.Push(&h, Element{nums[i][0], i, 0})
			if nums[i][0] > currentMaxNum {
				currentMaxNum = nums[i][0]
			}
		}
	}

	for h.Len() == len(nums) {
		top := heap.Pop(&h).(Element)
		if (rangeEnd - rangeStart) > (currentMaxNum - top.Val) {
			rangeStart = top.Val
			rangeEnd = currentMaxNum
		}
		top.EleInd++
		if top.EleInd < len(nums[top.ArrInd]) {
			top.Val = nums[top.ArrInd][top.EleInd]
			heap.Push(&h, top)
			if top.Val > currentMaxNum {
				currentMaxNum = top.Val
			}
		}
	}

	return []int{rangeStart, rangeEnd}
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

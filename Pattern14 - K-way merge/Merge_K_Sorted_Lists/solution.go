package mergeksortedlists

import "container/heap"

/*
* Merge K Sorted Lists
Given an array of ‘K’ sorted LinkedLists, merge them into one sorted list.

Input: L1=[2, 6, 8], L2=[3, 6, 7], L3=[1, 3, 4]
Output: [1, 2, 3, 3, 4, 6, 6, 7, 8]

Input: L1=[5, 8, 9], L2=[1, 7]
Output: [1, 5, 7, 8, 9]

ref: https://leetcode-cn.com/problems/merge-k-sorted-lists/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		finalHead *ListNode
		curHead   *ListNode
		h         MinHeap
	)
	heap.Init(&h)
	for _, head := range lists {
		if head != nil {
			heap.Push(&h, head)
		}
	}
	for h.Len() > 0 {
		small := heap.Pop(&h).(*ListNode)
		if finalHead == nil {
			finalHead = small
			curHead = small
		} else {
			curHead.Next = small
			curHead = curHead.Next
		}
		if small.Next != nil {
			heap.Push(&h, small.Next)
		}
	}

	return finalHead
}

type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

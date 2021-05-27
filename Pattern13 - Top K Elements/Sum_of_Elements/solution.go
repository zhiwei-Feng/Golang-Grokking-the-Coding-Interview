package sumofelements

import (
	"container/heap"
)

/*
* Given an array, find the sum of all numbers between the K1’th and K2’th smallest elements of that array.

Input: [1, 3, 12, 5, 15, 11], and K1=3, K2=6
Output: 23
Explanation: The 3rd smallest number is 5 and 6th smallest number 15. The sum of numbers coming
between 5 and 15 is 23 (11+12).

Input: [3, 5, 8, 7], and K1=1, K2=4
Output: 12
Explanation: The sum of the numbers between the 1st smallest number (3) and the 4th smallest
number (8) is 12 (5+7).
*/

func findSumOfElements(nums []int, k1, k2 int) int {
	var (
		minh       MinHeap
		sumOfElems int
	)

	heap.Init(&minh)
	for _, v := range nums {
		heap.Push(&minh, v)
	}

	// pop k1个元素
	for i := 0; i < k1; i++ {
		heap.Pop(&minh)
	}

	// sum k2-k1-1个元素
	for i := 0; i < k2-k1-1; i++ {
		sumOfElems += heap.Pop(&minh).(int)
	}

	return sumOfElems
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

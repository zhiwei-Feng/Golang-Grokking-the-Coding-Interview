package topkfrequentnumbers

import (
	"container/heap"
)

/*
Given an unsorted array of numbers, find the top ‘K’ frequently occurring numbers in it.

Input: [1, 3, 5, 12, 11, 12, 11], K = 2
Output: [12, 11]
Explanation: Both '11' and '12' apeared twice.

Input: [5, 12, 11, 3, 11], K = 2
Output: [11, 5] or [11, 12] or [11, 3]
Explanation: Only '11' appeared twice, all other numbers appeared once.

ref: https://leetcode-cn.com/problems/top-k-frequent-elements/
*/

type PairHeap [][2]int

func (ph PairHeap) Len() int            { return len(ph) }
func (ph PairHeap) Less(i, j int) bool  { return ph[i][1] < ph[j][1] }
func (ph PairHeap) Swap(i, j int)       { ph[i], ph[j] = ph[j], ph[i] }
func (ph *PairHeap) Push(x interface{}) { *ph = append(*ph, x.([2]int)) }
func (ph *PairHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	var (
		m  = make(map[int]int)
		ph PairHeap
	)

	for _, v := range nums {
		m[v]++
	}

	heap.Init(&ph)
	for key, value := range m {
		heap.Push(&ph, [2]int{key, value})
		if ph.Len() > k {
			heap.Pop(&ph)
		}
	}

	res := make([]int, 0, k)
	for _, v := range ph {
		res = append(res, v[0])
	}

	return res
}

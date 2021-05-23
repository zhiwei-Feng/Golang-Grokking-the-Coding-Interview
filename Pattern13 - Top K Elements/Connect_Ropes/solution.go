package connectropes

import "container/heap"

/*
Given ‘N’ ropes with different lengths, we need to connect these ropes into one big rope with minimum cost.
The cost of connecting two ropes is equal to the sum of their lengths.

Input: [1, 3, 11, 5]
Output: 33
Explanation: First connect 1+3(=4), then 4+5(=9), and then 9+11(=20). So the total cost is 33 (4+9+20)

Input: [3, 4, 5, 6]
Output: 36
Explanation: First connect 3+4(=7), then 5+6(=11), 7+11(=18). Total cost is 36 (7+11+18)

Input: [1, 3, 11, 5, 2]
Output: 42
Explanation: First connect 1+2(=3), then 3+3(=6), 6+5(=11), 11+11(=22). Total cost is 42 (3+6+11+22)

ref: https://leetcode-cn.com/problems/minimum-cost-to-connect-sticks/
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func connectSticks(sticks []int) int {
	var (
		res = 0
		h   IntHeap
	)

	for _, v := range sticks {
		h = append(h, v)
	}
	heap.Init(&h)

	for len(h) > 1 {
		x1 := heap.Pop(&h).(int)
		x2 := heap.Pop(&h).(int)
		sum := x1 + x2
		res += sum
		heap.Push(&h, sum)
	}

	return res
}

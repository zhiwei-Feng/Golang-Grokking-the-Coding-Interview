package maximumdistinctelements

import "container/heap"

/**********
*
* Given an array of numbers and a number ‘K’, we need to remove ‘K’ numbers from the array such that we are left with maximum distinct numbers.

Input: [7, 3, 5, 8, 5, 3, 3], and K=2
Output: 3
Explanation: We can remove two occurrences of 3 to be left with 3 distinct numbers [7, 3, 8], we have
to skip 5 because it is not distinct and occurred twice.
Another solution could be to remove one instance of '5' and '3' each to be left with three
distinct numbers [7, 5, 8], in this case, we have to skip 3 because it occurred twice.

Input: [3, 5, 12, 11, 12], and K=3
Output: 2
Explanation: We can remove one occurrence of 12, after which all numbers will become distinct. Then
we can delete any two numbers which will leave us 2 distinct numbers in the result.

Input: [1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 5], and K=2
Output: 3
Explanation: We can remove one occurrence of '4' to get three distinct numbers.

ref: https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals/
***********/

func findMaximumDistinctElements(nums []int, k int) int {
	//
	// minHeap存储多次出现的数及其出现次数
	//
	if len(nums) <= k {
		return 0
	}

	var (
		numHeap NumberHeap
		m       = make(map[int]int)
		res     = 0
	)

	for _, v := range nums {
		m[v]++
	}

	for key, val := range m {
		if val > 1 {
			numHeap = append(numHeap, Number{key, val})
		} else {
			res++
		}
	}
	heap.Init(&numHeap)

	for k > 0 && numHeap.Len() > 0 {
		top := heap.Pop(&numHeap).(Number)
		if top.Times-1 <= k {
			res++
		}
		k -= top.Times - 1
	}

	if k > 0 {
		res -= k
	}

	return res
}

type Number struct {
	Val   int
	Times int
}

type NumberHeap []Number

func (nh NumberHeap) Len() int            { return len(nh) }
func (nh NumberHeap) Less(i, j int) bool  { return nh[i].Times < nh[j].Times }
func (nh NumberHeap) Swap(i, j int)       { nh[i], nh[j] = nh[j], nh[i] }
func (nh *NumberHeap) Push(x interface{}) { *nh = append(*nh, x.(Number)) }
func (nh *NumberHeap) Pop() interface{} {
	old := *nh
	x := old[len(old)-1]
	*nh = old[:len(old)-1]
	return x
}

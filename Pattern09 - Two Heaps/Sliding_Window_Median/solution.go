package Sliding_Window_Median

import (
	"container/heap"
)

/*
Given an array of numbers and a number ‘k’, find the median of all the ‘k’ sized sub-arrays (or windows) of the array.

example1
Input: nums=[1, 2, -1, 3, 5], k = 2
Output: [1.5, 0.5, 1.0, 4.0]
Explanation: Lets consider all windows of size ‘2’:

[1, 2] -> median is 1.5
[2, -1] -> median is 0.5
[-1, 3] -> median is 1.0
[3, 5] -> median is 4.0

example2
Input: nums=[1, 2, -1, 3, 5], k = 3
Output: [1.0, 2.0, 3.0]
Explanation: Lets consider all windows of size ‘3’:

[1, 2, -1] -> median is 1.0
[2, -1, 3] -> median is 2.0
[-1, 3, 5] -> median is 3.0
*/

// 显然这是一个在滑动窗口中的过程中加入找中位数的算法。

type MaxHeap []int
type MinHeap []int

// impl sort.Interface
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// impl heap.Interface
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// impl Top
func (h MaxHeap) Top() int {
	return h[0]
}
func (h MinHeap) Top() int {
	return h[0]
}

type MedianFinder struct {
	Small *MaxHeap
	Big   *MinHeap
}

func Construct() MedianFinder {
	obj := MedianFinder{}
	obj.Small = &MaxHeap{}
	obj.Big = &MinHeap{}
	heap.Init(obj.Small)
	heap.Init(obj.Big)
	return obj
}

func (mf *MedianFinder) AddNum(x int) {

	if mf.Small.Len() == 0 || mf.Small.Top() >= x {
		heap.Push(mf.Small, x)
	} else {
		heap.Push(mf.Big, x)
	}

	if mf.Small.Len() > mf.Big.Len()+1 {
		heap.Push(mf.Big, heap.Pop(mf.Small))
	} else if mf.Small.Len() < mf.Big.Len() {
		heap.Push(mf.Small, heap.Pop(mf.Big))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.Small.Len() == mf.Big.Len() {
		return float64(mf.Small.Top()+mf.Big.Top()) / 2
	} else {
		return float64(mf.Small.Top())
	}
}

func reBalance(mf *MedianFinder) {
	if mf.Small.Len() > mf.Big.Len()+1 {
		heap.Push(mf.Big, heap.Pop(mf.Small))
	} else if mf.Small.Len() < mf.Big.Len() {
		heap.Push(mf.Small, heap.Pop(mf.Big))
	}
}

//
// 算法入口
//
func medianSlidingWindow(nums []int, k int) []float64 {
	if len(nums) == 0 {
		return []float64{}
	}
	var (
		windowStart = 0
		results     []float64
		mf          = Construct()
	)

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		mf.AddNum(nums[windowEnd])
		if windowEnd < k-1 {
			continue
		}
		results = append(results, mf.FindMedian())
		//remove from heap
		removeElementFromHeap(nums[windowStart], mf)
		// reBalance
		reBalance(&mf)
		windowStart++
	}

	return results
}

func removeElementFromHeap(x int, mf MedianFinder) {
	if x > mf.Small.Top() {
		big := *mf.Big
		for i, v := range big {
			if v == x {
				heap.Remove(mf.Big, i)
				break
			}
		}
	} else {
		small := *mf.Small
		for i, v := range small {
			if v == x {
				heap.Remove(mf.Small, i)
				break
			}
		}
	}
}

package Find_the_Median_of_a_Number_Stream

import (
	"container/heap"
)

/*
Design a class to calculate the median of a number stream. The class should have the following two methods:

1. insertNum(int num): stores the number in the class
2. findMedian(): returns the median of all numbers inserted in the class

If the count of numbers inserted in the class is even, the median will be the average of the middle two numbers.

example1:
1. insertNum(3)
2. insertNum(1)
3. findMedian() -> output: 2
4. insertNum(5)
5. findMedian() -> output: 3
6. insertNum(4)
7. findMedian() -> output: 3.5

ref: https://leetcode-cn.com/problems/find-median-from-data-stream/
*/

type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Top() int {
	return h[0]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) Top() int {
	return h[0]
}

type MedianFinder struct {
	Small *MaxHeap
	Big   *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	obj := MedianFinder{}
	obj.Small = &MaxHeap{}
	heap.Init(obj.Small)
	obj.Big = &MinHeap{}
	heap.Init(obj.Big)
	return obj
}

func (this *MedianFinder) AddNum(num int) {
	if this.Small.Len() == 0 || this.Small.Top() >= num {
		heap.Push(this.Small, num)
	} else {
		heap.Push(this.Big, num)
	}

	if this.Small.Len() > this.Big.Len()+1 {
		heap.Push(this.Big, heap.Pop(this.Small))
	} else if this.Small.Len() < this.Big.Len() {
		heap.Push(this.Small, heap.Pop(this.Big))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Small.Len() == this.Big.Len() {
		return float64(this.Small.Top()+this.Big.Top()) / 2
	} else {
		return float64(this.Small.Top())
	}
}

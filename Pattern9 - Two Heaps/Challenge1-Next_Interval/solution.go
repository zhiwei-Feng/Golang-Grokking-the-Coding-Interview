package Challenge1_Next_Interval

import "container/heap"

/*
Given an array of intervals, find the next interval of each interval.
In a list of intervals, for an interval ‘i’ its next interval ‘j’ will have the smallest ‘start’ greater than or equal to the ‘end’ of ‘i’.

Write a function to return an array containing indices of the next interval of each input interval.
If there is no next interval of a given interval, return -1. It is given that none of the intervals have the same start point.

Input: Intervals [[2,3], [3,4], [5,6]]
Output: [1, 2, -1]
Explanation: The next interval of [2,3] is [3,4] having index ‘1’. Similarly, the next interval of [3,4] is [5,6] having index ‘2’. There is no next interval for [5,6] hence we have ‘-1’.

Input: Intervals [[3,4], [1,5], [4,6]]
Output: [2, -1, -1]
Explanation: The next interval of [3,4] is [4,6] which has index ‘2’. There is no next interval for [1,5] and [4,6].
*/

func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 0 {
		return []int{}
	}
	var (
		minLInterval = MinLeftHeap{}
		minRInterval = MinRightHeap{}
		nextInterval = make([]int, len(intervals))
	)

	heap.Init(&minLInterval)
	heap.Init(&minRInterval)

	for i := 0; i < len(intervals); i++ {
		heap.Push(&minRInterval, Point{x: intervals[i][0], y: intervals[i][1], ind: i})
		heap.Push(&minLInterval, Point{x: intervals[i][0], y: intervals[i][1], ind: i})
	}

	for minRInterval.Len() != 0 {
		curPoint := heap.Pop(&minRInterval).(Point)
		for minLInterval.Len() != 0 && minLInterval[0].x < curPoint.y {
			heap.Pop(&minLInterval)
		}
		if minLInterval.Len() == 0 {
			nextInterval[curPoint.ind] = -1
		} else {
			nextInterval[curPoint.ind] = minLInterval[0].ind
		}
	}

	return nextInterval
}

type Point struct {
	x, y int
	ind  int
}

type MinRightHeap []Point
type MinLeftHeap []Point

func (m MinRightHeap) Len() int {
	return len(m)
}

func (m MinRightHeap) Less(i, j int) bool {
	return m[i].y < m[j].y
}

func (m MinRightHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinRightHeap) Push(x interface{}) {
	*m = append(*m, x.(Point))
}

func (m *MinRightHeap) Pop() interface{} {
	old := *m
	x := old[len(*m)-1]
	*m = old[:len(*m)-1]
	return x
}

func (m MinLeftHeap) Len() int {
	return len(m)
}

func (m MinLeftHeap) Less(i, j int) bool {
	return m[i].x < m[j].x
}

func (m MinLeftHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinLeftHeap) Push(x interface{}) {
	*m = append(*m, x.(Point))
}

func (m *MinLeftHeap) Pop() interface{} {
	old := *m
	x := old[len(*m)-1]
	*m = old[:len(*m)-1]
	return x
}

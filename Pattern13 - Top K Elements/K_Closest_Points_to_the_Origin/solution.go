package kclosestpointstotheorigin

import "container/heap"

/*
Given an array of points in a 2D2D plane, find ‘K’ closest points to the origin.

Input: points = [[1,2],[1,3]], K = 1
Output: [[1,2]]
Explanation: The Euclidean distance between (1, 2) and the origin is sqrt(5).
The Euclidean distance between (1, 3) and the origin is sqrt(10).
Since sqrt(5) < sqrt(10), therefore (1, 2) is closer to the origin.

Input: point = [[1, 3], [3, 4], [2, -1]], K = 2
Output: [[1, 3], [2, -1]]

ref: https://leetcode-cn.com/problems/k-closest-points-to-origin/
*/

type Point struct {
	x, y int
	dis  int
}

type PointHeap []Point

func (ph PointHeap) Len() int           { return len(ph) }
func (ph PointHeap) Less(i, j int) bool { return ph[i].dis > ph[j].dis }
func (ph PointHeap) Swap(i, j int)      { ph[i], ph[j] = ph[j], ph[i] }

func (ph *PointHeap) Push(x interface{}) { *ph = append(*ph, x.(Point)) }
func (ph *PointHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[:n-1]
	return x
}

func distance(x, y int) int {
	return x*x + y*y
}

func kClosest(points [][]int, k int) [][]int {
	var (
		ph  PointHeap
		i   = 0
		res = make([][]int, 0)
	)

	for ; i < k; i++ {
		ph = append(ph, Point{points[i][0], points[i][1], distance(points[i][0], points[i][1])})
	}
	heap.Init(&ph)
	for ; i < len(points); i++ {
		if distance(points[i][0], points[i][1]) < ph[0].dis {
			heap.Pop(&ph)
			heap.Push(&ph, Point{points[i][0], points[i][1], distance(points[i][0], points[i][1])})
		}
	}

	for _, v := range ph {
		res = append(res, []int{v.x, v.y})
	}
	return res
}

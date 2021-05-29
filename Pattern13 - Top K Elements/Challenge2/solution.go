package challenge2

import "container/heap"

/*
* Scheduling Tasks
You are given a list of tasks that need to be run, in any order, on a server.
Each task will take one CPU interval to execute but once a task has finished, it has a cooling period during which it can’t be run again.
If the cooling period for all tasks is ‘K’ intervals, find the minimum number of CPU intervals that the server needs to finish all tasks.

If at any time the server can’t execute any task then it must stay idle.

Input: [a, a, a, b, c, c], K=2
Output: 7
Explanation: a -> c -> b -> a -> c -> idle -> a

Input: [a, b, a], K=3
Output: 5
Explanation: a -> b -> idle -> idle -> a

ref: https://leetcode-cn.com/problems/task-scheduler/
*/

func leastInterval(tasks []byte, k int) int {
	var (
		freqMap       = make(map[byte]int)
		maxHeap       TaskHeap
		intervalCount = 0
	)

	for _, v := range tasks {
		freqMap[v]++
	}

	heap.Init(&maxHeap)
	for k, v := range freqMap {
		heap.Push(&maxHeap, Task{k, v})
	}

	for maxHeap.Len() > 0 {
		waitList := make([]Task, 0)
		n := k + 1
		// try to execute as many as 'k+1' tasks from the max-heap
		for ; n > 0 && maxHeap.Len() > 0; n-- {
			intervalCount++
			top := heap.Pop(&maxHeap).(Task)
			if top.Times > 1 {
				top.Times--
				waitList = append(waitList, top)
			}
		}

		// put all the waiting list back on the heap
		for _, it := range waitList {
			heap.Push(&maxHeap, it)
		}

		if maxHeap.Len() > 0 {
			// 如果maxHeap还有余，说明还有任务要执行
			// 此时的n表示经过这一轮任务执行后剩下的idle时间
			intervalCount += n
		}
	}

	return intervalCount
}

type Task struct {
	Val   byte
	Times int
}
type TaskHeap []Task

func (h TaskHeap) Len() int            { return len(h) }
func (h TaskHeap) Less(i, j int) bool  { return h[i].Times > h[j].Times }
func (h TaskHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x interface{}) { *h = append(*h, x.(Task)) }
func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

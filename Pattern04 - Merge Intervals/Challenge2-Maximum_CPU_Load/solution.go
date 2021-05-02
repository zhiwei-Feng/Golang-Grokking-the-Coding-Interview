package Challenge2_Maximum_CPU_Load

import (
	"container/heap"
	"sort"
)

/*
We are given a list of Jobs. Each job has a Start time, an End time, and a CPU load when it is running.
Our goal is to find the maximum CPU load at any time if all the jobs are running on the same machine.

Jobs: [[1,4,3], [2,5,4], [7,9,6]]
Output: 7
Explanation: Since [1,4,3] and [2,5,4] overlap, their maximum CPU load (3+4=7) will be when both the
jobs are running at the same time i.e., during the time interval (2,4).

Jobs: [[6,7,10], [2,4,11], [8,12,15]]
Output: 15
Explanation: None of the jobs overlap, therefore we will take the maximum load of any job which is 15.

Jobs: [[1,4,2], [2,4,1], [3,6,5]]
Output: 8
Explanation: Maximum CPU load will be 8 as all jobs overlap during the time interval [3,4].
*/

type Job struct {
	Start   int
	End     int
	CPULoad int
}

type minHeap []Job

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].End < h[j].End
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Job))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minHeap) Top() Job {
	old := *h
	return old[0]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func findMaxCPULoad(jobs []Job) int {
	if len(jobs) < 1 {
		return 0
	}
	sort.SliceStable(jobs, func(i, j int) bool {
		return jobs[i].Start < jobs[j].Start
	})

	h := &minHeap{}
	heap.Init(h)

	var (
		curLoad = 0
		maxLoad = 0
	)
	for _, j := range jobs {
		for h.Len() > 0 && h.Top().End <= j.Start {
			xInterface := heap.Pop(h)
			x := xInterface.(Job)
			curLoad -= x.CPULoad
		}
		heap.Push(h, j)
		curLoad += j.CPULoad
		maxLoad = max(maxLoad, curLoad)
	}

	return maxLoad
}

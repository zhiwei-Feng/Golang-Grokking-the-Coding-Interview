package Challenge1_Minimum_Meeting_Room

import (
	"container/heap"
	"sort"
)

/*
Given a list of intervals representing the start and end time of ‘N’ meetings,
find the minimum number of rooms required to hold all the meetings.

Meetings: [[1,4], [2,5], [7,9]]
Output: 2
Explanation: Since [1,4] and [2,5] overlap, we need two rooms to hold these two meetings. [7,9] can
occur in any of the two rooms later.

Meetings: [[6,7], [2,4], [8,12]]
Output: 1
Explanation: None of the meetings overlap, therefore we only need one room to hold all meetings.

Meetings: [[1,4], [2,3], [3,6]]
Output:2
Explanation: Since [1,4] overlaps with the other two meetings [2,3] and [3,6], we need two rooms to
hold all the meetings.

Meetings: [[4,5], [2,3], [2,4], [3,5]]
Output: 2
Explanation: We will need one room for [2,3] and [3,5], and another room for [2,4] and [4,5].
*/

type Meeting struct {
	Start int
	End   int
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

type minHeap []Meeting

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
	*h = append(*h, x.(Meeting))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minHeap) Top() Meeting {
	old := *h
	return old[0]
}

// 使用小根堆来做，不是常规的merge interval问题
func findMinimumMeetingRooms(meetings []Meeting) int {
	if len(meetings) < 1 {
		return 0
	}
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})

	h := &minHeap{}
	heap.Init(h)

	var count = 0
	for _, meet := range meetings {
		for h.Len() != 0 && h.Top().End <= meet.Start {
			heap.Pop(h)
		}
		heap.Push(h, meet)
		count = max(count, h.Len())
	}

	//var count = 0
	//var maxCount = 0
	//var start = meetings[0].Start
	//var end = meetings[0].End
	//for i := 0; i < len(meetings); i++ {
	//	if i == 0 {
	//		count++
	//		continue
	//	}
	//	if meetings[i].Start < end {
	//		// 交错
	//		count++
	//		start = max(start, meetings[i].Start)
	//		end = min(end, meetings[i].End)
	//	} else {
	//		maxCount = max(maxCount, count)
	//		count = 0
	//		start = meetings[i].Start
	//		end = meetings[i].End
	//	}
	//}

	return count
}

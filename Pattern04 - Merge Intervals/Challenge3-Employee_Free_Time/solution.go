package Challenge3_Employee_Free_Time

import (
	"container/heap"
)

/*
For ‘K’ employees, we are given a list of intervals representing the working hours of each employee.
Our goal is to find out if there is a free interval that is common to all employees.
You can assume that each list of employee working hours is sorted on the start time.

Input: Employee Working Hours=[[[1,3], [5,6]], [[2,3], [6,8]]]
Output: [3,5]
Explanation: Both the employees are free between [3,5].

Input: Employee Working Hours=[[[1,3], [9,12]], [[2,4]], [[6,8]]]
Output: [4,6], [8,9]
Explanation: All employees are free between [4,6] and [8,9].

Input: Employee Working Hours=[[[1,3]], [[2,4]], [[3,5], [7,9]]]
Output: [5,7]
Explanation: All employees are free between [5,7].
*/

type Interval struct {
	Start int
	End   int
}

type EmployeeInterval struct {
	interval      Interval
	employeeIndex int
	intervalIndex int
}

type minHeap []EmployeeInterval

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].interval.Start < h[j].interval.Start
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Top() EmployeeInterval {
	return h[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(EmployeeInterval))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findEmployeeFreeTime(schedule [][]Interval) []Interval {
	// 思路：利用小根堆来解决
	// 1. 维护一个小根堆保存每个employee当前start最小的interval
	// 2. 比较前两个最小start的interval的之间是否有free time
	// 3. pop掉最小start的interval，并把其employee的下一个interval加入到堆中.
	// 4. 直至结束.
	if len(schedule) < 1 {
		return []Interval{}
	}

	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < len(schedule); i++ {
		heap.Push(h, EmployeeInterval{
			interval:      schedule[i][0],
			employeeIndex: i,
			intervalIndex: 0,
		})
	}

	var (
		prevInterval = h.Top().interval
		freeTime     = make([]Interval, 0, len(schedule[0]))
	)
	for h.Len() > 0 {
		topInterval := heap.Pop(h).(EmployeeInterval)
		if prevInterval.End < topInterval.interval.Start {
			freeTime = append(freeTime, Interval{prevInterval.End, topInterval.interval.Start})
			prevInterval = topInterval.interval
		} else if prevInterval.End < topInterval.interval.End {
			prevInterval = topInterval.interval
		}

		topIntervalList := schedule[topInterval.employeeIndex]
		if topInterval.intervalIndex < len(topIntervalList)-1 {
			heap.Push(h, EmployeeInterval{
				interval:      topIntervalList[topInterval.intervalIndex+1],
				employeeIndex: topInterval.employeeIndex,
				intervalIndex: topInterval.intervalIndex + 1,
			})
		}
	}

	//var (
	//	allSchedule = schedule[0]
	//	freeTime    = make([]Interval, 0, len(schedule[0]))
	//)
	//// 1.将所有working hours 归到一个数组中
	//for i := 1; i < len(schedule); i++ {
	//	curEmployee := schedule[i]
	//	allSchedule = append(allSchedule, curEmployee...)
	//}
	//sort.SliceStable(allSchedule, func(i, j int) bool {
	//	return allSchedule[i].Start < allSchedule[j].Start
	//})
	//// 2.提出空隙区间
	//end := allSchedule[0].End
	//for _, sche := range allSchedule {
	//	if sche.Start <= end {
	//		// 重叠
	//		end = max(end, sche.End)
	//	} else {
	//		freeTime = append(freeTime, Interval{end, sche.Start})
	//		end = sche.End
	//	}
	//}
	//
	return freeTime
}

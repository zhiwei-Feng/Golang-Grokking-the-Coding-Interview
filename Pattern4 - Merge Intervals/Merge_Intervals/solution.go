package Merge_Intervals

import (
	"sort"
)

/*
Given a list of intervals,
merge all the overlapping intervals to produce a list that has only mutually exclusive intervals.

Intervals: [[1,4], [2,5], [7,9]]
Output: [[1,5], [7,9]]
Explanation: Since the first two intervals [1,4] and [2,5] overlap, we merged them into
one [1,5].

Intervals: [[6,7], [2,4], [5,9]]
Output: [[2,4], [5,9]]
Explanation: Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].

Intervals: [[1,4], [2,6], [3,5]]
Output: [[1,6]]
Explanation: Since all the given intervals overlap, we merged them into one.
*/

type Interval struct {
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

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	// sort
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start <= intervals[j].Start
	})
	var mergedList = make([]Interval, 0, len(intervals))
	var start = intervals[0].Start
	var end = intervals[0].End

	for _, currInterv := range intervals {
		if currInterv.Start <= end {
			end = max(end, currInterv.End)
		} else {
			mergedList = append(mergedList, Interval{start, end})
			start = currInterv.Start
			end = currInterv.End
		}
	}
	mergedList = append(mergedList, Interval{start, end})

	return mergedList
}

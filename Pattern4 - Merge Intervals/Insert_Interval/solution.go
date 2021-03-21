package Insert_Interval

/*
Given a list of non-overlapping intervals sorted by their start time,
insert a given interval at the correct position and merge all necessary intervals to produce a list that has only mutually exclusive intervals.

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,10]
Output: [[1,3], [4,12]]
Explanation: After insertion, since [4,10] overlaps with [5,7] & [8,12], we merged them into [4,12].

Input: Intervals=[[2,3],[5,7]], New Interval=[1,4]
Output: [[1,4], [5,7]]
Explanation: After insertion, since [1,4] overlaps with [2,3], we merged them into one [1,4].
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

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	var (
		mergedIntervals = make([]Interval, 0, len(intervals)+1)
		i               = 0
	)

	for i < len(intervals) && intervals[i].End < newInterval.Start {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		newInterval.Start = min(intervals[i].Start, newInterval.Start)
		newInterval.End = max(intervals[i].End, newInterval.End)
		i++
	}

	mergedIntervals = append(mergedIntervals, newInterval)

	for i < len(intervals) {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}

	return mergedIntervals
}

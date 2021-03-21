package Intervals_Intersection

/*
Given two lists of intervals, find the intersection of these two lists.
Each list consists of disjoint intervals sorted on their start time.

Input: arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]
Output: [2, 3], [5, 6], [7, 7]
Explanation: The output list contains the common intervals between the two lists.

Input: arr1=[[1, 3], [5, 7], [9, 12]], arr2=[[5, 10]]
Output: [5, 7], [9, 10]
Explanation: The output list contains the common intervals between the two lists.
*/

type Interval struct {
	Start int
	End   int
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func intersection(arr1, arr2 []Interval) []Interval {
	var (
		interList = make([]Interval, 0, min(len(arr1), len(arr2)))
		i, j      = 0, 0
	)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i].End < arr2[j].Start {
			i++
			continue
		}
		if arr2[j].End < arr1[i].Start {
			j++
			continue
		}
		// 相交
		start := max(arr1[i].Start, arr2[j].Start)
		end := min(arr1[i].End, arr2[j].End)
		interList = append(interList, Interval{start, end})

		if arr1[i].End < arr2[j].End {
			i++
		} else {
			j++
		}
	}

	return interList
}

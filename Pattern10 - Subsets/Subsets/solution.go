package Subsets

import "math"

/*
Given a set with distinct elements, find all of its distinct subsets.

Input: [1, 3]
Output: [], [1], [3], [1,3]

Input: [1, 5, 3]
Output: [], [1], [5], [3], [1,5], [1,3], [5,3], [1,5,3]

ref: https://leetcode-cn.com/problems/Subsets/
*/

func subsets(nums []int) [][]int {
	var results = make([][]int, 0, int(math.Pow(2, float64(len(nums)))))
	results = append(results, []int{})
	for _, currentNum := range nums {
		n := len(results)
		for i := 0; i < n; i++ {
			set := make([]int, len(results[i]))
			copy(set, results[i])
			set = append(set, currentNum)
			results = append(results, set)
		}
	}

	return results
}

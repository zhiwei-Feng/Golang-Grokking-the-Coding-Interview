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
	var results = make([][]int, int(math.Pow(2, float64(len(nums)))))
	results[0] = []int{}
	i := 1
	for _, currentNum := range nums {
		for _, subset := range results[:i] {
			newSubset := make([]int, len(subset)+1)
			copy(newSubset, subset)
			newSubset[len(subset)] = currentNum
			results[i] = newSubset
			i++
		}
	}

	return results
}

package Subsets_with_Duplicates

import (
	"math"
	"sort"
)

/*
Given a set of numbers that might contain duplicates, find all of its distinct subsets.

Input: [1, 3, 3]
Output: [], [1], [3], [1,3], [3,3], [1,3,3]

Input: [1, 5, 3, 3]
Output: [], [1], [5], [3], [1,5], [1,3], [5,3], [1,5,3], [3,3], [1,3,3], [3,3,5], [1,5,3,3]

ref: https://leetcode-cn.com/problems/subsets-ii/
*/

func subsetsWithDup(nums []int) [][]int {
	var (
		ans       = make([][]int, 0)
		prevStart = 0
		prevEnd   = 0
		prevNum   = math.MinInt8
	)

	ans = append(ans, []int{})
	prevEnd = 1

	// 1. sort nums
	sort.Ints(nums)

	for _, currentNum := range nums {
		prevAppend := ans[prevStart:prevEnd]
		prevStart = prevEnd
		if prevNum == math.MinInt8 || prevNum != currentNum {
			for _, subset := range ans[:prevEnd] {
				newSubset := make([]int, len(subset))
				copy(newSubset, subset)
				newSubset = append(newSubset, currentNum)
				ans = append(ans, newSubset)
				prevEnd++
			}

		} else {
			for _, subset := range prevAppend {
				// same operation as above
				newSubset := make([]int, len(subset))
				copy(newSubset, subset)
				newSubset = append(newSubset, currentNum)
				ans = append(ans, newSubset)
				prevEnd++
			}
		}
		prevNum = currentNum
	}

	return ans
}

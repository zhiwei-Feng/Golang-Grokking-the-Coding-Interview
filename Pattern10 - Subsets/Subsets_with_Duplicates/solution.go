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
		ans      = make([][]int, 0, int(math.Pow(2, float64(len(nums)))))
		startInd = 0
		endInd   = 0
	)
	sort.Ints(nums)
	ans = append(ans, []int{})

	for i := 0; i < len(nums); i++ {
		startInd = 0
		if i > 0 && nums[i] == nums[i-1] {
			startInd = endInd + 1
		}
		endInd = len(ans) - 1
		for _, subset := range ans[startInd : endInd+1] {
			newSubset := make([]int, len(subset), len(subset)+1)
			copy(newSubset, subset)
			newSubset = append(newSubset, nums[i])
			ans = append(ans, newSubset)
		}
	}

	return ans
}

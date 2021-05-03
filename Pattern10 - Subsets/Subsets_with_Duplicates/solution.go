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

//
// 1. 将nums排序
// 2. 对于currentNum==prevNum，我们只对上一次添加的子集进行新添操作
// 3. 原因是因为对于在上一次之前的子集，上一次已经进行了一次新添操作，当前轮再做一次操作会导致重复子集的出现，因此不再做操作。
//
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

package equalsubsetsumpartition

/*
* Equal Subset Sum Partition
Given a set of positive numbers, find if we can partition it into two subsets such that the sum of elements in both subsets is equal.

Input: {1, 2, 3, 4}
Output: True
Explanation: The given set can be partitioned into two subsets with equal sum: {1, 4} & {2, 3}

Input: {1, 1, 3, 4, 7}
Output: True
Explanation: The given set can be partitioned into two subsets with equal sum: {1, 3, 4} & {1, 7}

Input: {2, 3, 4, 6}
Output: False
Explanation: The given set cannot be partitioned into two subsets with equal sum.

ref: https://leetcode-cn.com/problems/partition-equal-subset-sum
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)

	for s := 1; s <= target; s++ {
		// corner
		dp[s] = nums[0] == s
	}

	for i := 1; i < len(nums); i++ {
		for j := target; j > 0; j-- {
			if !dp[j] && j>=nums[i] {
				dp[j] = dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}

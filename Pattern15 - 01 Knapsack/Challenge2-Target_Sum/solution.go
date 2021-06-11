package Challenge2_Target_Sum

/*
* Target Sum
You are given a set of positive numbers and a target sum ‘S’.
Each number should be assigned either a ‘+’ or ‘-’ sign.
We need to find the total ways to assign symbols to make the sum of the numbers equal to the target ‘S’.

Input: {1, 1, 2, 3}, S=1
Output: 3
Explanation: The given set has '3' ways to make a sum of '1': {+1-1-2+3} & {-1+1-2+3} & {+1+1+2-3}

Input: {1, 2, 7, 1}, S=9
Output: 2
Explanation: The given set has '2' ways to make a sum of '9': {+1+2+7-1} & {-1+2+7+1}

ref: https://leetcode-cn.com/problems/target-sum/
*/

//
// idea:
// Sum(s1) - Sum(s2) = target
// Sum(s1) + Sum(s2) = Sum(num)
// 其中s1中都是+号的数，s2中都是-号的数
// 所以Sum(s1) = (target + Sum(num)) / 2
//
func findTargetSumWays(nums []int, target int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	if totalSum < target || (totalSum+target)%2 != 0 {
		return 0
	}

	target = (totalSum + target) / 2

	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i <= target; i++ {
		if nums[0] == i {
			dp[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for s := 1; s <= target; s++ {
			dp[i][s] = dp[i-1][s]
			if s >= nums[i] {
				dp[i][s] += dp[i-1][s-nums[i]]
			}
		}
	}

	return dp[n-1][target]
}

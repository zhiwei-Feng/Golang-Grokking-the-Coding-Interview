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

func findTargetSumWays(nums []int, target int) int {
	return recursive(nums, 0, target, 0)
}

func recursive(nums []int, curInd int, target int, sum int) int {
	if curInd == len(nums) {
		if sum == target {
			return 1
		} else {
			return 0
		}
	}

	count1 := recursive(nums, curInd+1, target, sum+nums[curInd])
	count2 := recursive(nums, curInd+1, target, sum-nums[curInd])

	return count1 + count2
}

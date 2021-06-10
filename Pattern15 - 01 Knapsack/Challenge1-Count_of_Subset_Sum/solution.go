package challenge1countofsubsetsum

/********************************
Count of Subset Sum
Given a set of positive numbers, find the total number of subsets whose sum is equal to a given number ‘S’.

Input: {1, 1, 2, 3}, S=4
Output: 3
The given set has '3' subsets whose sum is '4': {1, 1, 2}, {1, 3}, {1, 3}
Note that we have two similar sets {1, 3}, because we have two '1' in our input.

Input: {1, 2, 7, 1, 5}, S=9
Output: 3
The given set has '3' subsets whose sum is '9': {2, 7}, {1, 7, 1}, {1, 2, 1, 5}
*/

func countSubsets(nums []int, sum int) int {
	//dp[i]表示和为j的个数
	n := len(nums)
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 1; i <= sum; i++ {
		if nums[0] == i {
			dp[i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for s := sum; s >= 0; s-- {
			if s >= nums[i] {
				dp[s] += dp[s-nums[i]]
			}
		}
	}

	return dp[sum]
}

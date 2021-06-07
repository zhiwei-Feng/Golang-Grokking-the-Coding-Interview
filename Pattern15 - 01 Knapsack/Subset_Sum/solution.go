package subsetsum

/*
* Given a set of positive numbers, determine if a subset exists whose sum is equal to a given number ‘S’.

Input: {1, 2, 3, 7}, S=6
Output: True
The given set has a subset whose sum is '6': {1, 2, 3}

Input: {1, 2, 7, 1, 5}, S=10
Output: True
The given set has a subset whose sum is '10': {1, 2, 7}

Input: {1, 3, 4, 8}, S=6
Output: False
The given set does not have any subset whose sum is equal to '6'.

*/

func canPartition(num []int, sum int) bool {
	// dp[i][j] 前i个数是否可以得到j
	n := len(num)
	var dp = make([]bool, sum+1)
	dp[0] = true
	for s := 1; s <= sum; s++ {
		dp[s] = num[0] == s
	}

	for i := 1; i < n; i++ {
		for s := sum; s >= 0; s-- {
			if !dp[s] && s >= num[i] {
				dp[s] = dp[s-num[i]]
			}
		}
	}

	return dp[sum]
}

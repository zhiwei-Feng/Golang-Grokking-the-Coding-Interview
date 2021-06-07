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
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= sum; i++ {
		if num[0] == i {
			dp[0][i] = true
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= num[i] && dp[i-1][j-num[i]] {
				dp[i][j] = true
			}
		}
	}

	return dp[n-1][sum]
}

package minimumsubsetsumdifference

/* ----------------
Given a set of positive numbers, partition the set into two subsets with minimum difference between their subset sums.

Input: {1, 2, 3, 9}
Output: 3
Explanation: We can partition the given set into two subsets where minimum absolute difference
between the sum of numbers is '3'. Following are the two subsets: {1, 2, 3} & {9}.

Input: {1, 2, 7, 1, 5}
Output: 0
Explanation: We can partition the given set into two subsets where minimum absolute difference
between the sum of number is '0'. Following are the two subsets: {1, 2, 5} & {7, 1}.

Input: {1, 3, 100, 4}
Output: 92
Explanation: We can partition the given set into two subsets where minimum absolute difference
between the sum of numbers is '92'. Here are the two subsets: {1, 3, 4} & {100}.
*/

func canPartition(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum / 2
	// dp[i][j]表示前i个数的和是否等于j
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	for i := 0; i <= target; i++ {
		dp[0][i] = nums[0] == i
	}

	for i := 1; i < len(nums); i++ {
		for s := 1; s <= target; s++ {
			if dp[i-1][s] {
				dp[i][s] = true
			} else if s >= nums[i] {
				dp[i][s] = dp[i-1][s-nums[i]]
			}
		}
	}

	sum1 := 0
	for i := target; i >= 0; i-- {
		if dp[len(nums)-1][i] {
			sum1 = i
			break
		}
	}
	sum2 := sum - sum1
	return abs(sum1 - sum2)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

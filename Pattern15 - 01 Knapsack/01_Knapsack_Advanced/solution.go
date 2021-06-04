package knapsackadvanced

/*
improved solution for 0/1 knapsack
*/

func solveKnapsack(profits, weights []int, capacity int) int {
	//
	// 进一步优化只使用一维数组来存储dp值
	// 需要注意遍历capacity的时候需要从大到小来遍历
	//
	if capacity <= 0 || len(profits) == 0 || len(profits) != len(weights) {
		return 0
	}

	n := len(profits)
	dp := make([]int, capacity+1)
	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := capacity; c >= 0; c-- {
			prof1, prof2 := 0, 0
			if weights[i] <= c {
				prof1 = profits[i] + dp[c-weights[i]]
			}
			prof2 = dp[c]
			dp[c] = max(prof1, prof2)
		}
	}

	return dp[capacity]
}

func solveKnapsack2(profits, weights []int, capacity int) int {
	//
	// because dp[i][j] mainly depend on dp[i-1][j] and dp[i-1][j-weights[i]]
	// so we just need two array to store dp value
	// and dp[i%2][j] depend on dp[(i-1)%2][j] and dp[(i-1)%2][j]
	//
	if capacity <= 0 || len(profits) == 0 || len(profits) != len(weights) {
		return 0
	}

	n := len(profits)

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c], dp[1][c] = profits[0], profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 0; c <= capacity; c++ {
			prof1, prof2 := 0, 0
			if weights[i] <= c {
				prof1 = profits[i] + dp[(i-1)%2][c-weights[i]]
			}
			prof2 = dp[(i-1)%2][c]
			dp[i%2][c] = max(prof1, prof2)
		}
	}

	return dp[(n-1)%2][capacity]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

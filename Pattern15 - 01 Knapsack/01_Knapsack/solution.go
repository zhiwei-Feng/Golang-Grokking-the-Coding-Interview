package knapsack

/*
* 0/1 Knapsack
Given two integer arrays to represent weights and profits of ‘N’ items,
we need to find a subset of these items which will give us maximum profit such that their cumulative weight is not more than a given number ‘C.’
Each item can only be selected once, which means either we put an item in the knapsack or we skip it.
*/

func solveKnapsack(profits, weights []int, capacity int) int {
	if capacity <= 0 || len(profits) == 0 || len(profits) != len(weights) {
		return 0
	}

	n := len(profits)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 1; c <= capacity; c++ {
			profit1, profit2 := 0, 0
			if weights[i] <= c {
				profit1 = profits[i] + dp[i-1][c-weights[i]]
			}
			profit2 = dp[i-1][c]
			dp[i][c] = max(profit1, profit2)
		}
	}

	return dp[n-1][capacity]
}

func solveKnapsackByMemorization(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := 0; i < len(profits); i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 0; i < len(profits); i++ {
		for j := 0; j < capacity+1; j++ {
			dp[i][j] = -1
		}
	}
	return knapsackRecursiveWithMemory(dp, profits, weights, capacity, 0)
}

func knapsackRecursiveWithMemory(dp [][]int, profits, weights []int, capacity int, currentInd int) int {
	if capacity <= 0 || currentInd >= len(profits) {
		return 0
	}

	if dp[currentInd][capacity] != -1 {
		return dp[currentInd][capacity]
	}

	profit1 := 0
	if weights[currentInd] <= capacity {
		profit1 = profits[currentInd] + knapsackRecursive(profits, weights, capacity-weights[currentInd], currentInd+1)
	}
	profit2 := knapsackRecursive(profits, weights, capacity, currentInd+1)

	dp[currentInd][capacity] = max(profit1, profit2)
	return dp[currentInd][capacity]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solveKnapsackByRecursive(profits, weights []int, capacity int) int {
	return knapsackRecursive(profits, weights, capacity, 0)
}

func knapsackRecursive(profits, weights []int, capacity int, currentInd int) int {
	if capacity <= 0 || currentInd >= len(profits) {
		return 0
	}

	profit1 := 0
	if weights[currentInd] <= capacity {
		profit1 = profits[currentInd] + knapsackRecursive(profits, weights, capacity-weights[currentInd], currentInd+1)
	}
	profit2 := knapsackRecursive(profits, weights, capacity, currentInd+1)

	if profit1 > profit2 {
		return profit1
	}
	return profit2
}

package Maximum_Sum_Subarray_of_Size_K

/*
Given an array of positive numbers and a positive number k,
find the maximum sum of any contiguous subarray of size k.

Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].

Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findMaxSumSubArray(k int, arr []int) int {
	// use sliding window approach
	var (
		maxSum      = 0
		windowSum   = 0
		windowStart = 0
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

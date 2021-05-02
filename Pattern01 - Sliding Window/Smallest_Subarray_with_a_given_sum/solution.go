package Smallest_Subarray_with_a_given_sum

/*
Given an array of positive numbers and a positive number ‘S,’
find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
Return 0 if no such subarray exists.

Input: [2, 1, 5, 2, 3, 2], S=7
Output: 2
Explanation: The smallest subarray with a sum great than or equal to '7' is [5, 2].

Input: [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].

Input: [3, 4, 1, 1, 6], S=8
Output: 3
Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] or [1, 1, 6].
*/

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func findMinSubArray(S int, arr []int) int {
	var (
		windowStart = 0
		windowSum   = 0
		smallest    = -1
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowSum < S {
			continue
		}
		for windowSum >= S {
			if smallest == -1 {
				smallest = windowEnd - windowStart + 1
			} else {
				smallest = min(smallest, windowEnd-windowStart+1)
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	if smallest == -1 {
		return 0
	}
	return smallest
}

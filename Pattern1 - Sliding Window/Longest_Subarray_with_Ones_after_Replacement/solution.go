package Longest_Subarray_with_Ones_after_Replacement

/*
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s,
find the length of the longest contiguous subarray having all 1s.

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/

func findLength(arr []int, k int) int {
	var (
		maxLength   = 0
		windowStart = 0
		zeroNum     = 0
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 0 {
			zeroNum++
		}

		for zeroNum > k {
			if arr[windowStart] == 0 {
				zeroNum--
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

package bitonicarraymaximum

/*
Find the maximum value in a given Bitonic array.
An array is considered bitonic if it is monotonically increasing and then monotonically decreasing.
Monotonically increasing or decreasing means that for any index i in the array arr[i] != arr[i+1].

Input: [1, 3, 8, 12, 4, 2]
Output: 12
Explanation: The maximum number in the input bitonic array is '12'.

Input: [3, 8, 3, 1]
Output: 8

Input: [1, 3, 8, 12]
Output: 12

Input: [10, 9, 8]
Output: 10

ref: https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
ref: https://leetcode-cn.com/problems/find-peak-element/
*/

func findMax(arr []int) int {
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start < end {
		mid := start + (end-start)/2
		if arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return arr[start]
}

package challenge3rotationcount

/*
Given an array of numbers which is sorted in ascending order and is rotated ‘k’ times around a pivot, find ‘k’.

You can assume that the array does not have any duplicates.

Input: [10, 15, 1, 3, 8]
Output: 2
Explanation: The array has been rotated 2 times.

Input: [4, 5, 7, 9, 10, -1, 2]
Output: 5
Explanation: The array has been rotated 5 times.

Input: [1, 3, 8, 10]
Output: 0
Explanation: The array has been not been rotated.

ref: https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
*/

func countRotations(arr []int) int {
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start < end {
		mid := start + (end-start)/2
		if mid < end && arr[mid] > arr[mid+1] {
			return mid + 1
		}
		if mid > start && arr[mid-1] > arr[mid] {
			return mid
		}

		if arr[start] < arr[mid] {
			// left part sorted, so pivot is right part
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return 0
}

package challenge2searchinarotatedarray

/*
Given an array of numbers which is sorted in ascending order and also rotated by some arbitrary number,
find if a given ‘key’ is present in it.

Write a function to return the index of the ‘key’ in the rotated array. If the ‘key’ is not present, return -1.
You can assume that the given array does not have any duplicates.

Input: [10, 15, 1, 3, 8], key = 15
Output: 1
Explanation: '15' is present in the array at index '1'.

Input: [4, 5, 7, 9, 10, -1, 2], key = 10
Output: 4
Explanation: '10' is present in the array at index '4'.

ref: https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
*/

func search(nums []int, target int) int {
	var breakInd = -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			breakInd = i
			break
		}
	}
	if breakInd == -1 {
		breakInd = len(nums) - 1
	}

	keyInd := binarySearch(nums, target, 0, breakInd)
	if keyInd == -1 {
		keyInd = binarySearch(nums, target, breakInd+1, len(nums)-1)
	}

	return keyInd
}

func binarySearch(arr []int, key int, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		}
		if key < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

package Ceiling_of_a_Number

/*
Given an array of numbers sorted in an ascending order, find the ceiling of a given number ‘key’.
The ceiling of the ‘key’ will be the smallest element in the given array greater than or equal to the ‘key’.

Write a function to return the index of the ceiling of the ‘key’. If there isn’t any ceiling return -1.

Input: [4, 6, 10], key = 6
Output: 1
Explanation: The smallest number greater than or equal to '6' is '6' having index '1'.

Input: [1, 3, 8, 10, 15], key = 12
Output: 4
Explanation: The smallest number greater than or equal to '12' is '15' having index '4'.

Input: [4, 6, 10], key = 17
Output: -1
Explanation: There is no number greater than or equal to '17' in the given array.

Input: [4, 6, 10], key = -1
Output: 0
Explanation: The smallest number greater than or equal to '-1' is '4' having index '0'.
*/

func searchCeilingOfANumber(arr []int, key int) int {
	if arr[len(arr)-1] < key {
		return -1
	}
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] < key {
			start = mid + 1
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			return mid
		}
	}

	return start
}

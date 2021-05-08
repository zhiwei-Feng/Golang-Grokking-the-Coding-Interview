package orderagnosticbinarysearch

/*
Given a sorted array of numbers, find if a given number ‘key’ is present in the array.
Though we know that the array is sorted, we don’t know if it’s sorted in ascending or descending order.
You should assume that the array can have duplicates.

Write a function to return the index of the ‘key’ if it is present in the array, otherwise return -1.

Input: [4, 6, 10], key = 10
Output: 2

Input: [1, 2, 3, 4, 5, 6, 7], key = 5
Output: 4

Input: [10, 6, 4], key = 10
Output: 0

Input: [10, 6, 4], key = 4
Output: 2
*/

func search(arr []int, key int) int {
	var (
		start = 0
		end   = len(arr) - 1
		isAsc = arr[start] < arr[end]
	)

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == key {
			return mid
		}

		if isAsc {
			if key < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if key > arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

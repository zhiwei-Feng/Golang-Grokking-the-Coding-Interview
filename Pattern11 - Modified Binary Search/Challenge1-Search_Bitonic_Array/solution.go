package challenge1searchbitonicarray

/*
Given a Bitonic array, find if a given ‘key’ is present in it.
An array is considered bitonic if it is monotonically increasing and then monotonically decreasing.
Monotonically increasing or decreasing means that for any index i in the array arr[i] != arr[i+1].

Write a function to return the index of the ‘key’. If the ‘key’ is not present, return -1.

Input: [1, 3, 8, 4, 3], key=4
Output: 3

Input: [3, 8, 3, 1], key=8
Output: 1

Input: [1, 3, 8, 12], key=12
Output: 3

Input: [10, 9, 8], key=10
Output: 0
*/

func search(arr []int, key int) int {
	peekInd := findMax(arr)
	keyInd := binarySearch(arr, key, 0, peekInd)
	if keyInd == -1 {
		keyInd = binarySearch(arr, key, peekInd+1, len(arr)-1)
	}

	return keyInd
}

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

	return start
}

func binarySearch(arr []int, key int, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		}
		if arr[start] < arr[end] {
			if key < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if key < arr[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

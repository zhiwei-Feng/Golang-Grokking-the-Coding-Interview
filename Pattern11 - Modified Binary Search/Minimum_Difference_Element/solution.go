package minimumdifferenceelement

/*
Given an array of numbers sorted in ascending order, find the element in the array that has the minimum difference with the given ‘key’.

Input: [4, 6, 10], key = 7
Output: 6
Explanation: The difference between the key '7' and '6' is minimum than any other number in the array

Input: [4, 6, 10], key = 4
Output: 4

Input: [1, 3, 8, 10, 15], key = 12
Output: 10

Input: [4, 6, 10], key = 17
Output: 10
*/

func searchMinDiffElement(arr []int, key int) int {
	if key < arr[0] {
		return arr[0]
	}
	if key > arr[len(arr)-1] {
		return arr[len(arr)-1]
	}
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return arr[mid]
		}
		if key < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// end+1 == start, arr[end]<key<arr[start]
	if arr[start]-key < key-arr[end] {
		return arr[start]
	}
	return arr[end]
}

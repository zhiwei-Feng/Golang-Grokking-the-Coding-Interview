package numberrange

/*
Given an array of numbers sorted in ascending order, find the range of a given number ‘key’.
The range of the ‘key’ will be the first and last position of the ‘key’ in the array.

Write a function to return the range of the ‘key’. If the ‘key’ is not present return [-1, -1].

Input: [4, 6, 6, 6, 9], key = 6
Output: [1, 3]

Input: [1, 3, 8, 10, 15], key = 10
Output: [3, 3]

Input: [1, 3, 8, 10, 15], key = 12
Output: [-1, -1]
*/

func searchRange(nums []int, target int) []int {
	var res = [2]int{-1, -1}

	res[0] = search(nums, target, true)
	if res[0] != -1 {
		res[1] = search(nums, target, false)
	}

	return res[:]
}

func search(arr []int, target int, findLeftCorner bool)  int {
	var (
		start = 0
		end   = len(arr) - 1
		index = -1
	)

	for start <= end {
		mid := start + (end-start)/2
		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			index = mid
			if findLeftCorner {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return index
}

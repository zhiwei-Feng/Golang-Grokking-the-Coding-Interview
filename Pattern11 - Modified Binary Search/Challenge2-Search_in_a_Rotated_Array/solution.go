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
	var (
		start = 0
		end   = len(nums) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}

		if nums[start] <= nums[mid] {
			//左侧有序
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

/*
Similar Problem
How do we search in a sorted and rotated array that also has duplicates?

Input: [3, 7, 3, 3, 3], key = 7
Output: 1
Explanation: '7' is present in the array at index '1'.

ref: https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
*/
func search2(nums []int, target int) int {
	var (
		start = 0
		end   = len(nums) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}

		if nums[start] == nums[mid] && nums[mid] == nums[end] {
			start++
			end--
			continue
		}

		if nums[start] <= nums[mid] {
			//左侧有序
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

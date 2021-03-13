package Challenge3_Minimum_Window_Sort

import "math"

/*
Given an array, find the length of the smallest subarray in it which when sorted will sort the whole array.

Input: [1, 2, 5, 3, 7, 10, 9, 12]
Output: 5
Explanation: We need to sort only the subarray [5, 3, 7, 10, 9] to make the whole array sorted

Input: [1, 3, 2, 0, -1, 7, 10]
Output: 5
Explanation: We need to sort only the subarray [1, 3, 2, 0, -1] to make the whole array sorted

Input: [1, 2, 3]
Output: 0
Explanation: The array is already sorted

Input: [3, 2, 1]
Output: 3
Explanation: The whole array needs to be sorted.
*/

func sort(arr []int) int {
	var (
		left     = 0
		right    = len(arr) - 1
		minOfSub = math.MaxInt8
		maxOfSub = math.MinInt8
	)

	// 1. 从左找到第一个无序的元素
	for left < len(arr)-1 && arr[left] <= arr[left+1] {
		left++
	}
	if left == len(arr)-1 {
		return 0
	}

	// 2. 从右找到第一个无序的元素
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	// 3. 找到子数组的最小和最大值
	for i := left; i <= right; i++ {
		if arr[i] > maxOfSub {
			maxOfSub = arr[i]
		}
		if arr[i] < minOfSub {
			minOfSub = arr[i]
		}
	}

	// 4. 往左拓展和往右拓展
	for left > 0 && arr[left-1] > minOfSub {
		left--
	}
	for right < len(arr)-1 && arr[right+1] < maxOfSub {
		right++
	}

	return right - left + 1
}

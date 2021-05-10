package numberrange

// import "sync"

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
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	var (
		start = 0
		end   = len(nums) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			// 从当前位置找开始点和结尾点
			// var waitG sync.WaitGroup
			// waitG.Add(1)
			// waitG.Add(1)
			// go func(ind int) {
			// 	for ind >= 0 && nums[ind] == target {
			// 		start = ind
			// 		ind--
			// 	}
			// 	waitG.Done()
			// }(mid)

			// go func(ind int) {
			// 	for ind < len(nums) && nums[ind] == target {
			// 		end = ind
			// 		ind++
			// 	}
			// 	waitG.Done()
			// }(mid)

			// waitG.Wait()
			for i := mid; i >= 0 && nums[i] == target; i-- {
				start = i
			}
			for i := mid; i < len(nums) && nums[i] == target; i++ {
				end = i
			}
			return []int{start, end}
		}
	}

	return []int{-1, -1}
}

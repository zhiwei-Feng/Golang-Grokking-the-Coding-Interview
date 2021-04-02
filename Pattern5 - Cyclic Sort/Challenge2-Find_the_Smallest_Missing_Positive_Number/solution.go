package Challenge2_Find_the_Smallest_Missing_Positive_Number

/*
Given an unsorted array containing numbers, find the smallest missing positive number in it.

Input: [-3, 1, 5, 4, 2]
Output: 3
Explanation: The smallest missing positive number is '3'

Input: [3, -2, 0, 1, 2]
Output: 4

Input: [3, 2, 5, 1]
Output: 4
*/

func findNumber(nums []int) int {
	var i = 0
	for i < len(nums) {
		// 只对 1..n的值进行cyclic sort，确保其在正确的位置上
		// 比如1在index 0上，这样排序后，遍历数组，只需要判断第一个value!=index+1的位置就好
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

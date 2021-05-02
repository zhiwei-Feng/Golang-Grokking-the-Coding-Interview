package Challenge1_Find_the_Corrupt_Pair

import "fmt"

/*
We are given an unsorted array containing ‘n’ numbers taken from the range 1 to ‘n’.
The array originally contained all the numbers from 1 to ‘n’, but due to a data error,
one of the numbers got duplicated which also resulted in one number going missing. Find both these numbers.

Input: [3, 1, 2, 5, 2]
Output: [2, 4]
Explanation: '2' is duplicated and '4' is missing.

Input: [3, 1, 2, 3, 6, 4]
Output: [3, 5]
Explanation: '3' is duplicated and '5' is missing.
*/

func findNumbers(nums []int) []int {
	var i = 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	fmt.Println(nums)

	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return []int{nums[i], i + 1}
		}
	}
	return []int{-1, -1}
}

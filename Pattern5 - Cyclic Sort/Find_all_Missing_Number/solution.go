package Find_all_Missing_Number

/*
We are given an unsorted array containing numbers taken from the range 1 to ‘n’. The array can have duplicates,
which means some numbers will be missing. Find all those missing numbers.

Input: [2, 3, 1, 8, 2, 3, 5, 1]
Output: 4, 6, 7
Explanation: The array should have all numbers from 1 to 8, due to duplicates 4, 6, and 7 are missing.

Input: [2, 4, 1, 2]
Output: 3

Input: [2, 3, 2, 1]
Output: 4
*/

func findNumbers(nums []int) []int {
	var i = 0
	var results = make([]int, 0, len(nums))
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			results = append(results, i+1)
		}
	}
	return results
}

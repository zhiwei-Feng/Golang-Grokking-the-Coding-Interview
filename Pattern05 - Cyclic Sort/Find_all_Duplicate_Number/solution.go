package Find_all_Duplicate_Number

/*
We are given an unsorted array containing ‘n’ numbers taken from the range 1 to ‘n’.
The array has some numbers appearing twice,
find all these duplicate numbers without using any extra space.

Input: [3, 4, 4, 5, 5]
Output: [4, 5]

Input: [5, 4, 7, 2, 3, 5, 3]
Output: [3, 5]
*/

func findAllDuplicateNumbers(nums []int) []int {
	var i = 0
	var duplicated = make([]int, 0, len(nums))
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
			duplicated = append(duplicated, nums[i])
		}
	}

	return duplicated
}

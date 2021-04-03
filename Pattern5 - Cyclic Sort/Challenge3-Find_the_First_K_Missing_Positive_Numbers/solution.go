package Challenge3_Find_the_First_K_Missing_Positive_Numbers

/*
Given an unsorted array containing numbers and a number ‘k’, find the first ‘k’ missing positive numbers in the array.

Input: [3, -1, 4, 5, 5], k=3
Output: [1, 2, 6]
Explanation: The smallest missing positive numbers are 1, 2 and 6.

Input: [2, 3, 4], k=3
Output: [1, 5, 6]
Explanation: The smallest missing positive numbers are 1, 5 and 6.

Input: [-2, -3, 4], k=2
Output: [1, 2]
Explanation: The smallest missing positive numbers are 1 and 2.
*/

func findNumbers(nums []int, k int) []int {
	var (
		missingNums  = make([]int, 0, len(nums))
		i            = 0
		alreadyExist = make(map[int]int)
	)
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	for i = 0; i < len(nums) && len(missingNums) < k; i++ {
		if nums[i] != i+1 {
			missingNums = append(missingNums, i+1)
			alreadyExist[nums[i]] = 1
		}
	}

	for len(missingNums) < k {
		if _, ok := alreadyExist[i+1]; !ok {
			missingNums = append(missingNums, i+1)
		}
		i++
	}

	return missingNums
}

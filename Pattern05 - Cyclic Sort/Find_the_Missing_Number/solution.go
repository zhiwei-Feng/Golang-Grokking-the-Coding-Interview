package Find_the_Missing_Number

/*
We are given an array containing ‘n’ distinct numbers taken from the range 0 to ‘n’.
Since the array has only ‘n’ numbers out of the total ‘n+1’ numbers, find the missing number.

Input: [4, 0, 3, 1]
Output: 2

Input: [8, 3, 5, 2, 4, 6, 0, 1]
Output: 7
*/

func findMissingNumber(nums []int) int {
	// version1 存储已经见过的数字 -> time:O(n), space:O(n)
	//var (
	//	see = make([]int, len(nums)+1)
	//)
	//for _, v := range nums {
	//	see[v] = 1
	//}
	//for i, v := range see {
	//	if v == 0 {
	//		return i
	//	}
	//}
	//return -1
	// version2 -> time:O(n)
	var i = 0
	for i < len(nums) {
		if nums[i] < len(nums) && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

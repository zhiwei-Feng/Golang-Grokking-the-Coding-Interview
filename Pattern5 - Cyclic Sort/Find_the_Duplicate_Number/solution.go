package Find_the_Duplicate_Number

/*
We are given an unsorted array containing ‘n+1’ numbers taken from the range 1 to ‘n’.
The array has only one duplicate but it can be repeated multiple times.
Find that duplicate number without using any extra space. You are, however, allowed to modify the input array.

Input: [1, 4, 4, 3, 2]
Output: 4

Input: [2, 1, 3, 3, 5, 4]
Output: 3

Input: [2, 4, 1, 4, 4]
Output: 4
*/

func findDuplicateNumber(nums []int) int {
	// version1 not best
	//for i < len(nums) {
	//	j := nums[i] - 1
	//	if nums[i] != nums[j] {
	//		nums[i], nums[j] = nums[j], nums[i]
	//	} else {
	//		i++
	//	}
	//}
	//for i = 0; i < len(nums); i++ {
	//	if nums[i] != i+1 {
	//		return nums[i]
	//	}
	//}

	// version2 时间最快，但是会修改原数组
	//for i < len(nums) {
	//	j := nums[i] - 1
	//	// duplicate的情况满足两个条件
	//	// 1. 当前位置不满足条件 nums[i]!=i+1
	//	// 2. 需要交换的位置满足条件 nums[i] == nums[j]
	//	if nums[i] != i+1 {
	//		if nums[i] != nums[j] {
	//			nums[i], nums[j] = nums[j], nums[i]
	//		} else {
	//			return nums[i]
	//		}
	//	} else {
	//		i++
	//	}
	//}

	// version3 不修改原数组，利用fast-slow pointer求解
	// 因为有重复的树则必会有回路，且回路的起点就是重复的数
	var (
		fast      = 0
		slow      = 0
		circleLen = 0
	)

	fast = nums[nums[fast]]
	slow = nums[slow]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	current := nums[slow]
	circleLen++
	for current != slow {
		current = nums[current]
		circleLen++
	}
	slow = 0
	fast = 0
	for i := 0; i < circleLen; i++ {
		fast = nums[fast]
	}
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}

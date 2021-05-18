package twosinglenumbers

/*
In a non-empty array of numbers, every number appears exactly twice except two numbers that appear only once.
Find the two numbers that appear only once.

Input: [1, 4, 2, 1, 3, 5, 6, 2, 3, 5]
Output: [4, 6]

Input: [2, 1, 3, 2]
Output: [1, 3]

ref: https://leetcode-cn.com/problems/single-number-iii/
*/

func singleNumber(nums []int) []int {
	n1xn2 := 0
	for _, v := range nums {
		n1xn2 ^= v
	}

	rightMostSetBit := 1
	for rightMostSetBit&n1xn2 == 0 {
		rightMostSetBit = rightMostSetBit << 1
	}

	num1 := 0
	num2 := 0
	for _, v := range nums {
		if v&rightMostSetBit != 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}

	return []int{num1, num2}
}

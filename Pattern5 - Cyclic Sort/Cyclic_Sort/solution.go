package Cyclic_Sort

/*
We are given an array containing ‘n’ objects.
Each object, when created, was assigned a unique number from 1 to ‘n’ based on their creation sequence

Write a function to sort the objects in-place on their creation sequence number in O(n)O(n) and without any extra space.

Input: [3, 1, 5, 4, 2]
Output: [1, 2, 3, 4, 5]

Input: [2, 6, 4, 3, 1, 5]
Output: [1, 2, 3, 4, 5, 6]

Input: [1, 5, 6, 4, 3, 2]
Output: [1, 2, 3, 4, 5, 6]
*/

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

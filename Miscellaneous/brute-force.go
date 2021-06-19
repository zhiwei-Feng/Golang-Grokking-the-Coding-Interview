// Package Miscellaneous
/*
Kth Smallest Number:
Given an unsorted array of numbers, find Kth smallest number in it.
Please note that it is the Kth smallest number in the sorted order, not the Kth distinct element.

Input: [1, 5, 12, 2, 11, 5], K = 3
Output: 5
Explanation: The 3rd smallest number is '5', as the first two smaller numbers are [1, 2].

Input: [1, 5, 12, 2, 11, 5], K = 4
Output: 5
Explanation: The 4th smallest number is '5', as the first three smaller numbers are
[1, 2, 5].

Input: [5, 12, 11, -1, 12], K = 3
Output: 11
Explanation: The 3rd smallest number is '11', as the first two small numbers are [5, -1].
*/
package Miscellaneous

import "sort"

//
// time complexity: O(nlogn)
// space complexity: O(1)/O(n)
//
func findKthSmallestNumber(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	sort.Ints(nums)
	return nums[k-1]
}

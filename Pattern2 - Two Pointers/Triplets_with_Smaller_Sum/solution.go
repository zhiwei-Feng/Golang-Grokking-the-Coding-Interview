package Triplets_with_Smaller_Sum

import "sort"

/*
Given an array arr of unsorted numbers and a target sum,
count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices.
Write a function to return the count of such triplets.

Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]

Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]
*/

func searchTriplets(arr []int, target int) int {
	var (
		count = 0
	)

	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			currentSum := arr[i] + arr[left] + arr[right]
			//如果currentSum>=target, currentSum需要减小因此right--
			//如果currentSum<target, 首先left不变的情况下，right--直到right=left+1都是满足条件的
			//其次left++的情况下,不一定满足条件，但也需要遍历

			if currentSum >= target {
				right--
			} else {
				count += right - left
				left++
			}
		}
	}

	return count
}

package Triplet_Sum_Close_to_Target

import (
	"math"
	"sort"
)

/*
Given an array of unsorted numbers and a target number,
find a triplet in the array whose sum is as close to the target number as possible, return the sum of the triplet.
If there are more than one such triplet return the sum of the triplet with the smallest sum.

Input: [-2, 0, 1, 2], target=2
Output: 1
Explanation: The triplet [-2, 1, 2] has the closest sum to the target.

Input: [-3, -1, 1, 2], target=1
Output: 0
Explanation: The triplet [-3, 1, 2] has the closest sum to the target.

Input: [1, 0, 1, 1], target=100
Output: 3
Explanation: The triplet [1, 1, 1] has the closest sum to the target.
*/

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func searchTriplet(arr []int, targetSum int) int {
	sort.Ints(arr)
	smallestDifference := math.MaxInt8
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]
			if targetDiff == 0 {
				return targetSum - targetDiff
			}

			if abs(targetDiff) < abs(smallestDifference) {
				smallestDifference = targetDiff
			}

			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return targetSum - smallestDifference
}

package Triplet_Sum_to_Zero

import "sort"

/*
Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.
*/

func searchTriplets(arr []int) [][3]int {
	// 1. sort array
	sort.Ints(arr)
	triplets := make([][3]int, 0)

	// 2. 把左侧的值当前target，在右侧剩余arr中运行Pair_with_Target_Sum算法
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			// 这里之所以可以跳过，是因为后面的查找操作，会把剩余arr中所有的triplets找到
			continue
		}
		triplets = searchPair(arr, -arr[i], i+1, triplets)
	}

	return triplets
}

func searchPair(arr []int, targetNum int, left int, triplets [][3]int) [][3]int {
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetNum {
			triplets = append(triplets, [3]int{-targetNum, arr[left], arr[right]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if currentSum > targetNum {
			right--
		} else {
			left++
		}
	}
	return triplets
}

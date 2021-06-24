package Miscellaneous

import (
	"sort"
)

//
// coverage time complexity: O(n), worst O(n)
// space complexity: O(n)
//
func findKthSmallestNumberUsingMedianOfMedians(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	return findPivotRecursive(nums, k, 0, len(nums)-1)
}

func findPivotRecursive(nums []int, k, start, end int) int {
	p := partitionMom(nums, start, end)
	if p == k-1 {
		return nums[p]
	}
	if p > k-1 {
		return findPivotRecursive(nums, k, start, p-1)
	}
	return findPivotRecursive(nums, k, p+1, end)
}

func partitionMom(nums []int, low, high int) int {
	if low == high {
		return low
	}

	median := medianOfMedians(nums, low, high)
	// median as pivot
	// swap high and median
	for i := low; i < high; i++ {
		if nums[i] == median {
			nums[i], nums[high] = nums[high], nums[i]
			break
		}
	}

	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func medianOfMedians(nums []int, low, high int) int {
	n := high - low + 1
	if n < 5 {
		return nums[low]
	}

	numOfPartitions := n / 5
	medians := make([]int, numOfPartitions)
	for i := 0; i < numOfPartitions; i++ {
		partStart := low + i*5
		sort.Ints(nums[partStart : partStart+5])
		medians[i] = nums[partStart+2]
	}

	return partitionMom(medians, 0, numOfPartitions-1)
}

package Miscellaneous

import (
	"math/rand"
	"time"
)

//
// coverage time complexity: O(n), worst O(n^2)
// space complexity: O(n)
//
func findKthSmallestNumberUsingRandomQuickSort(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	return recursiveFindPivot(nums, k, 0, len(nums)-1)
}

func recursiveFindPivot(nums []int, k, start, end int) int {
	p := partitionRandomized(nums, start, end)
	if p == k-1 {
		return nums[p]
	}
	if p > k-1 {
		return recursiveFindPivot(nums, k, start, p-1)
	}
	return recursiveFindPivot(nums, k, p+1, end)
}

func partitionRandomized(nums []int, low, high int) int {
	if low == high {
		return low
	}

	rand.Seed(time.Now().Unix())
	pivot_ind := low + rand.Intn(high-low)
	nums[pivot_ind], nums[high] = nums[high], nums[pivot_ind]

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

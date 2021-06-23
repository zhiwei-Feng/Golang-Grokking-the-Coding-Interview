package Miscellaneous

//
// coverage time complexity: O(n), worst O(n^2)
// space complexity: O(n)
//
func findKthSmallestNumberUsingQuickSort(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	return partitionRec(nums, k, 0, len(nums)-1)
}

func partitionRec(nums []int, k, start, end int) int {
	p := partition(nums, start, end)
	if p == k-1 {
		return nums[p]
	}
	if p > k-1 {
		return partitionRec(nums, k, start, p-1)
	}
	return partitionRec(nums, k, p+1, end)
}

func partition(nums []int, low, high int) int {
	if low == high {
		return low
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

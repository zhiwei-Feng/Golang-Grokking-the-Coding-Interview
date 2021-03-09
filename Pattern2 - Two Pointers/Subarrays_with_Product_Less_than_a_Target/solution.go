package Subarrays_with_Product_Less_than_a_Target

/*
Given an array with positive numbers and a target number,
find all of its contiguous subarrays whose product is less than the target number.

Input: [2, 5, 3, 10], target=30
Output: [2], [5], [2, 5], [3], [5, 3], [10]
Explanation: There are six contiguous subarrays whose product is less than the target.

Input: [8, 2, 6, 5], target=50
Output: [8], [2], [8, 2], [6], [2, 6], [5], [6, 5]
Explanation: There are seven contiguous subarrays whose product is less than the target.
*/

func findSubarrays(arr []int, target int) [][]int {
	// 需要用上sliding window pattern
	var (
		windowStart = 0
		windowProd  = 1
		subarrs     = make([][]int, 0)
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowProd *= arr[windowEnd]
		// ===solution 1===
		//if arr[windowEnd] < target {
		//	subarrs = append(subarrs, []int{arr[windowEnd]})
		//}
		//
		//for windowProd >= target {
		//	windowProd /= arr[windowStart]
		//	windowStart++
		//}
		//if windowStart!=windowEnd{
		//	subarrs = append(subarrs, arr[windowStart:windowEnd+1])
		//}
		// ===end===

		// solution 2
		for windowProd >= target {
			windowProd /= arr[windowStart]
			windowStart++
		}

		tempList := make([]int, 0)
		for right := windowEnd; right >= windowStart; right-- {
			tempList = append([]int{arr[right]}, tempList...)
			subarrs = append(subarrs, tempList)
		}
	}

	return subarrs
}

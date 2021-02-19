package introduction

func findAveragesOfSubArraysBySlidingWindow(K int, arr []int) []float64 {
	results := make([]float64, len(arr)-K+1)
	windowSum := 0
	windowStart := 0
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		// 把下个元素加上
		windowSum += arr[windowEnd]
		// 滑动窗口，特别地，如果没有达到K个则不滑动
		if windowEnd >= K-1 {
			// 当前windowSum计算了整个窗口的和
			results[windowStart] = float64(windowSum) / float64(K)
			// 当前窗口计算完后，移动窗口需要先减去原来窗口头部的元素
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return results
}

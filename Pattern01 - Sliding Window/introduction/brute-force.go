package introduction

func findAveragesOfSubArraysByBruteForce(K int, arr []int) []float64 {
	results := make([]float64, len(arr)-K+1)
	for i := 0; i <= len(arr)-K; i++ {
		// 计算下一组K个元素的和
		sum := 0
		for j := i; j < i+K; j++ {
			sum += arr[j]
		}
		// 计算平均
		results[i] = float64(sum) / float64(K)
	}
	return results
}

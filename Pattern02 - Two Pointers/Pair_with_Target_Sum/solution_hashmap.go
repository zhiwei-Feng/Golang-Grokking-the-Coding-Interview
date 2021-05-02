package Pair_with_Target_Sum

func searchByHashMap(arr []int, targetSum int) (int, int) {
	var record = make(map[int]int) // int -> index
	for i := 0; i < len(arr); i++ {
		if _, ok := record[targetSum-arr[i]]; ok {
			return record[targetSum-arr[i]], i
		} else {
			record[arr[i]] = i
		}
	}

	return -1, -1
}

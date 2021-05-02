package Reverse_Level_Order_Traversal

func method1() {
	result := make([][]int, 0)
	for i := 0; i < 100; i++ {
		item := make([]int, 0)
		for j := 0; j < 10; j++ {
			item = append(item, i*j)
		}
		join := [][]int{item}
		result = append(join, result...)
	}
}

func method2() {
	result := make([][]int, 0)
	for i := 0; i < 100; i++ {
		item := make([]int, 0)
		for j := 0; j < 10; j++ {
			item = append(item, i*j)
		}
		result = append(result, item)
	}
	tmp := result
	result = make([][]int, 0)
	for i := len(tmp) - 1; i >= 0; i-- {
		result = append(result, tmp[i])
	}
}

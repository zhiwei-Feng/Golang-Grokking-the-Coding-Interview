package Zigzag_Traversal

func method1() {
	var input = make([]int, 0, 100)
	for i := 0; i < len(input); i++ {
		input = append(input, i)
	}
	// method1
	for i := 0; i < len(input)/2; i++ {
		j := len(input) - 1 - i
		input[i], input[j] = input[j], input[i]
	}
}

func method2() {
	var input = make([]int, 0, 100)
	for i := 0; i < len(input); i++ {
		input = append(input, i)
	}
	// method2
	tmp := input
	input = make([]int, 0, 100)
	for i := len(tmp) - 1; i >= 0; i-- {
		input = append(input, tmp[i])
	}
}

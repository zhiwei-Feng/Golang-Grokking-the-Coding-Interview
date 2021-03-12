package Challenge1_Quadruple_Sum_to_Target

import (
	"sort"
)

/*
Given an array of unsorted numbers and a target number,
find all unique quadruplets in it, whose sum is equal to the target number.

Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.

Input: [2, 0, -1, 1, -2, 2], target=2
Output: [-2, 0, 2, 2], [-1, 0, 1, 2]
Explanation: Both the quadruplets add up to the target.
*/

func searchQuadruplets(arr []int, target int) [][4]int {
	sort.Ints(arr)
	quadruplets := make([][4]int, 0, 10)
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			quadruplets = searchPair(arr, j+1, target, arr[i], arr[j], quadruplets)
		}
	}
	return quadruplets
}

func searchPair(arr []int, left int, target int, first, second int, quadruplets [][4]int) [][4]int {
	// assume arr is sorted
	var (
		i, j = left, len(arr) - 1
	)
	for i < j {
		current := arr[i] + arr[j] + first + second
		if current == target {
			quadruplets = append(quadruplets, [4]int{first, second, arr[i], arr[j]})
			i++
			j--
			for i < j && arr[i] == arr[i-1] {
				i++
			}
			for i < j && arr[j] == arr[j+1] {
				j--
			}
		} else if current > target {
			j--
		} else {
			i++
		}
	}

	return quadruplets
}

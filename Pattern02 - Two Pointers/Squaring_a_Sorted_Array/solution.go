package Squaring_a_Sorted_Array

/*
Given a sorted array, create a new array containing squares of all the numbers of the input array in the sorted order.

Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]

Input: [-3, -1, 0, 1, 2]
Output: [0, 1, 1, 4, 9]
*/

func square(x int) int {
	return x * x
}

func makeSquares(arr []int) []int {
	var (
		left                = 0
		right               = len(arr) - 1
		results             = make([]int, len(arr))
		currentHighestIndex = len(arr) - 1
	)

	for left <= right {
		leftSquare := square(arr[left])
		rightSquare := square(arr[right])
		if leftSquare < rightSquare {
			results[currentHighestIndex] = rightSquare
			right--
		} else {
			results[currentHighestIndex] = leftSquare
			left++
		}
		currentHighestIndex--
	}

	return results
}

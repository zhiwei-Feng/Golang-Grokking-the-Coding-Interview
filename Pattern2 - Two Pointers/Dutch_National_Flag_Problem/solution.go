package Dutch_National_Flag_Problem

/*
Given an array containing 0s, 1s and 2s, sort the array in-place. You should treat numbers of the array as objects,
hence, we can’t count 0s, 1s, and 2s to recreate the array.

The flag of the Netherlands consists of three colors: red, white and blue;
and since our input array also consists of three different numbers that is why it is called Dutch National Flag problem.

Input: [1, 0, 2, 1, 0]
Output: [0 0 1 1 2]

Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]
*/

func sort(arr []int) {
	var (
		low  = 0
		high = len(arr) - 1
	)

	for i := 0; i <= high; {
		if arr[i] == 0 {
			arr[i], arr[low] = arr[low], arr[i]
			i++
			low++
		} else if arr[i] == 1 {
			i++
		} else {
			arr[i], arr[high] = arr[high], arr[i]
			// 注意不能i++，因为有可能换过来的是0
			high--
		}
	}
}

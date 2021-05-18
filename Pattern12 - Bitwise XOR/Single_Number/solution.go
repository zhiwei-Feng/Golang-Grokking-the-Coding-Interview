package singlenumber

/*
In a non-empty array of integers, every number appears twice except for one, find that single number.

Input: 1, 4, 2, 1, 3, 2, 3
Output: 4

Input: 7, 9, 7
Output: 9
*/

func singleNumber(nums []int) int {
	x := 0
	for _, v := range nums {
        x ^= v
    }

	return x
}

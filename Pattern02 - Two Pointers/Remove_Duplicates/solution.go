package Remove_Duplicates

/*
Given an array of sorted numbers, remove all duplicates from it.
You should not use any extra space; after removing the duplicates in-place return the length of the subarray
that has no duplicate in it.


Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].


Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].
*/

func remove(arr []int) int {
	// 因为有序，所以重复数字是连着的
	newInd := 0
	for curInd := 1; curInd < len(arr); curInd++ {
		if arr[curInd] == arr[newInd] {
			continue
		}
		newInd++
		arr[newInd] = arr[curInd]
	}
	return newInd + 1
}

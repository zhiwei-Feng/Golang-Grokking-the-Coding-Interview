package searchinasortedinfinitearray

import (
	"math"
)

/*
Given an infinite sorted array (or an array with unknown size), find if a given number ‘key’ is present in the array.
Write a function to return the index of the ‘key’ if it is present in the array, otherwise return -1.

Since it is not possible to define an array with infinite (unknown) size,
you will be provided with an interface ArrayReader to read elements of the array. ArrayReader.get(index) will return the number at index;
if the array’s size is smaller than the index, it will return Integer.MAX_VALUE.

Input: [4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30], key = 16
Output: 6
Explanation: The key is present at index '6' in the array.

Input: [4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30], key = 11
Output: -1
Explanation: The key is not present in the array.

Input: [1, 3, 8, 10, 15], key = 15
Output: 4
Explanation: The key is present at index '4' in the array.

Input: [1, 3, 8, 10, 15], key = 200
Output: -1
Explanation: The key is not present in the array.

ref: https://leetcode-cn.com/problems/search-in-a-sorted-array-of-unknown-size/
*/

type ArrayReader struct {
	arr []int
}

func (reader *ArrayReader) get(index int) int {
	if index >= len(reader.arr) {
		return math.MaxInt32
	}

	return reader.arr[index]
}

func search(reader ArrayReader, target int) int {
	var (
		start = 0
		end   = 1
	)

	for reader.get(end) < target {
		start, end = end+1, end+(end-start+1)*2
	}

	return binarySearch(reader, target, start, end)
}

func binarySearch(reader ArrayReader, target, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if reader.get(mid) == target {
			return mid
		}
		if target < reader.get(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

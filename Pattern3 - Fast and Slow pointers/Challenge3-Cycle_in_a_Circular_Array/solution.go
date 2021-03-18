package Challenge3_Cycle_in_a_Circular_Array

/*
We are given an array containing positive and negative numbers. Suppose the array contains a number ‘M’ at a particular index.
Now, if ‘M’ is positive we will move forward ‘M’ indices and if ‘M’ is negative move backwards ‘M’ indices.
You should assume that the array is circular which means two things:
1. If, while moving forward, we reach the end of the array, we will jump to the first element to continue the movement.
2. If, while moving backward, we reach the beginning of the array, we will jump to the last element to continue the movement.

Write a method to determine if the array has a cycle.
The cycle should have more than one element and should follow one direction which means the cycle should not contain both forward and backward movements.

Input: [1, 2, -1, 2, 2]
Output: true
Explanation: The array has a cycle among indices: 0 -> 1 -> 3 -> 0

Input: [2, 2, -1, 2]
Output: true
Explanation: The array has a cycle among indices: 1 -> 3 -> 1

Input: [2, 1, -1, -2]
Output: false
Explanation: The array does not have any cycle.
*/

func findNextIndex(arr []int, isForward bool, currentIndex int) int {
	var direction = arr[currentIndex] >= 0
	if isForward != direction {
		return -1
	}

	nextIndex := (currentIndex + arr[currentIndex] + len(arr)) % len(arr)
	if nextIndex == currentIndex {
		nextIndex = -1
	}
	return nextIndex
}

func loopExists(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		isForward := arr[i] >= 0
		slow, fast := i, i

		// 先进行一次move
		slow = findNextIndex(arr, isForward, slow)
		fast = findNextIndex(arr, isForward, fast)
		if fast != -1 {
			fast = findNextIndex(arr, isForward, fast)
		}
		for slow != -1 && fast != -1 && fast != slow {
			slow = findNextIndex(arr, isForward, slow)
			fast = findNextIndex(arr, isForward, fast)
			if fast != -1 {
				fast = findNextIndex(arr, isForward, fast)
			}
		}
		// 追逐直到汇聚的时候，如果fast和slow的方向都不变且没有单元素的环则说明符合要求
		if slow != -1 && fast == slow {
			return true
		}
	}

	return false
}

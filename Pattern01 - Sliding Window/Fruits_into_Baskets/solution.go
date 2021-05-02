package Fruits_into_Baskets

/*
Given an array of characters where each character represents a fruit tree,
you are given two baskets, and your goal is to put maximum number of fruits in each basket.
The only restriction is that each basket can have only one type of fruit.
You can start with any tree, but you can’t skip a tree once you have started.
You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

Write a function to return the maximum number of fruits in both baskets.

Input: Fruit=['A', 'B', 'C', 'A', 'C']
Output: 3
Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']

Input: Fruit=['A', 'B', 'C', 'B', 'B', 'C']
Output: 5
Explanation: We can put 3 'B' in one basket and two 'C' in the other basket.
This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']
*/

// 这一题其实算是一个变种的Longest Substring with K distinct characters问题
func findLength(arr []byte) int { // byte means char
	var (
		maxLength   = 0
		windowStart = 0
		distMap     = make(map[byte]int)
	)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		distMap[arr[windowEnd]]++

		for len(distMap) > 2 {
			delChar := arr[windowStart]
			distMap[delChar]--
			if distMap[delChar] == 0 {
				delete(distMap, delChar)
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

package Longest_Substring_with_K_Distinct_Characters

/*
Given a string, find the length of the longest substring in it with no more than K distinct characters.

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".

Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findLength(str string, k int) int {
	var (
		maxLength   = 0
		windowStart = 0
		distinctMap = make(map[uint8]int) // record each character appear times of window
	)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		curChar := str[windowEnd]
		distinctMap[curChar]++

		for len(distinctMap) > k {
			delChar := str[windowStart]
			distinctMap[delChar]--
			if distinctMap[delChar] == 0 {
				delete(distinctMap, delChar)
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

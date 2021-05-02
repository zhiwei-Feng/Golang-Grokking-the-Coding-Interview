package Longest_Substring_with_Same_Letters_after_Replacement

/*
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’ letters with any letter,
find the length of the longest substring having the same letters after replacement.

Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
*/

func findLength(str string, k int) int {
	var (
		maxLength    = 0
		windowStart  = 0
		charMap      = make(map[byte]int) // store the times of each characters appear
		MaxRepeatNum = 0
	)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		curChar := str[windowEnd]
		charMap[curChar]++
		MaxRepeatNum = max(charMap[curChar], MaxRepeatNum)

		for windowEnd-windowStart+1-MaxRepeatNum > k {
			delChar := str[windowStart]
			charMap[delChar]--
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

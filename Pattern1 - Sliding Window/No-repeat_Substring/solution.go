package No_repeat_Substring

/*
Given a string, find the length of the longest substring, which has no repeating characters.

Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".

Input: String="abbbb"
Output: 2
Explanation: The longest substring without any repeating characters is "ab".

Input: String="abccde"
Output: 3
Explanation: Longest substrings without any repeating characters are "abc" & "cde".
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// solution1: the same pattern as Longest_Substring_with_K_Distinct_Characters
//var (
//		maxLength = 0
//		windStart = 0
//		existMap  = make(map[byte]int)
//	)
//
//	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
//		existMap[str[windowEnd]]++
//
//		for !isNoRepeat(existMap) {
//			delChar := str[windStart]
//			existMap[delChar]--
//			if existMap[delChar] == 0 {
//				delete(existMap, delChar)
//			}
//			windStart++
//		}
//
//		maxLength = max(maxLength, windowEnd-windStart+1)
//	}
//
//	return maxLength
func findLength(str string) int {
	var (
		maxLength   = 0
		windowStart = 0
		lastIndex   = make(map[byte]int)
	)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		cur := str[windowEnd]
		if _, ok := lastIndex[cur]; ok {
			windowStart = max(windowStart, lastIndex[cur]+1) // key code
		}
		lastIndex[cur] = windowEnd
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

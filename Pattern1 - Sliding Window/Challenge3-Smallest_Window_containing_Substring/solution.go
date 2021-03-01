package Challenge3_Smallest_Window_containing_Substring

/*
Given a string and a pattern, find the smallest substring in the given string which has all the characters of the given pattern.

Input: String="aabdec", Pattern="abc"
Output: "abdec"
Explanation: The smallest substring having all characters of the pattern is "abdec"

Input: String="abdbca", Pattern="abc"
Output: "bca"
Explanation: The smallest substring having all characters of the pattern is "bca".

Input: String="adcad", Pattern="abc"
Output: ""
Explanation: No substring in the given string has all characters of the pattern.
*/

func findSubstring(str, pattern string) string {
	var (
		windowStart = 0
		matched     = 0
		minLength   = len(str) + 1
		subStrStart = 0
		patternMap  = make(map[byte]int)
	)
	for i := 0; i < len(pattern); i++ {
		patternMap[pattern[i]]++
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		wright := str[windowEnd]
		if _, ok := patternMap[wright]; ok {
			patternMap[wright]--
			if patternMap[wright] >= 0 { // key point 1, 不再是==0，而是>=0
				matched++
			}
		}

		for matched == len(patternMap) { // 如果当前window包含pattern所有的字符，则从左收缩窗口至不完全包含状态
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStrStart = windowStart
			}

			wleft := str[windowStart]
			if _, ok := patternMap[wleft]; ok {
				// 注意这里仅当count==0才matched--，以为重复出现的字符会导致patternMap中的值>1，
				// 但是matched的值只减过一次，因此在上边进行matched++的部分，必须修改==0为>=0，
				// 不然会出现窗口右扩的时候，pattern字符进行一次patternMap[wright]--后，count依旧>0，无法使matched++
				// 但实际上该字符已经出现了，是满足条件的，这就产生了bug.
				if patternMap[wleft] == 0 {
					matched--
				}
				patternMap[wleft]++
			}
			windowStart++
		}
	}

	if minLength > len(str) {
		return ""
	} else {
		return str[subStrStart : subStrStart+minLength]
	}
}

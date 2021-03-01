package Challenge2_String_Anagrams

/*
Given a string and a pattern, find all anagrams of the pattern in the given string.

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
*/

func findStringAnagrams(str, pattern string) []int {
	var (
		matchedStartIndex = make([]int, 0, 10)
		windowStart       = 0
		patternMap        = make(map[byte]int)
		matched           = 0
	)

	// 1. convert pattern string to pattern map
	for i := 0; i < len(pattern); i++ {
		patternMap[pattern[i]]++
	}

	// 2. sliding window process
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		wright := str[windowEnd]
		if _, ok := patternMap[wright]; ok {
			patternMap[wright]--
			if patternMap[wright] == 0 {
				matched++
			}
		}

		if matched == len(patternMap) {
			matchedStartIndex = append(matchedStartIndex, windowStart)
		}

		if windowEnd >= len(patternMap)-1 {
			wleft := str[windowStart]
			if _, ok := patternMap[wleft]; ok {
				if patternMap[wleft] == 0 {
					matched--
				}
				patternMap[wleft]++
			}
			windowStart++
		}
	}

	return matchedStartIndex
}

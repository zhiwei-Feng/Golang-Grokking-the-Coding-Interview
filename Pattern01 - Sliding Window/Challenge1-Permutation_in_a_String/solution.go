package Challenge1_Permutation_in_a_String

/*
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string.
For example, “abc” has the following six permutations:
abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters, it will have n!n! permutations.

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
*/

// 最初的solution
//func isSamePattern(s1, s2 map[byte]int) bool {
//	if len(s1) != len(s2) {
//		return false
//	}
//
//	for k, v := range s1 {
//		if v != s2[k] {
//			return false
//		}
//	}
//
//	return true
//}
//
//func findPermutation(str string, pattern string) bool {
//	var (
//		windowStart = 0
//		charMap     = make(map[byte]int)
//		patternMap  = make(map[byte]int)
//	)
//
//	// convert pattern to patternMap
//	for i := 0; i < len(pattern); i++ {
//		patternMap[pattern[i]]++
//	}
//
//	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
//		// 记录当前的char
//		charMap[str[windowEnd]]++
//
//		// 若是window过小则跳过
//		if windowEnd < len(pattern)-1 {
//			continue
//		}
//
//		// 判断
//		if isSamePattern(charMap, patternMap) {
//			return true
//		}
//
//		// 后处理
//		delChar := str[windowStart]
//		charMap[delChar]--
//		if charMap[delChar] == 0 {
//			delete(charMap, delChar)
//		}
//		windowStart++
//	}
//
//	return false
//}

// better solution
func findPermutation(str, pattern string) bool {
	var (
		matched     = 0
		windowStart = 0
		patternMap  = make(map[byte]int)
	)

	// convert pattern to patternMap
	for i := 0; i < len(pattern); i++ {
		patternMap[pattern[i]]++
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		current := str[windowEnd]

		// 如果当前字符存在map中，则次数--，如果次数为0，则matched++. 说明该字符完全匹配完毕
		if _, ok := patternMap[current]; ok {
			patternMap[current]--
			if patternMap[current] == 0 {
				matched++
			}
		}

		// 验证完当前字符，判断是否完全与pattern匹配
		if matched == len(patternMap) {
			return true
		}

		// 如果window大小超过了pattern，说明需要收缩window
		if windowEnd >= len(pattern)-1 {
			delChar := str[windowStart]
			if _, ok := patternMap[delChar]; ok {
				if patternMap[delChar] == 0 {
					matched--
				}
				patternMap[delChar]++
			}
			windowStart++
		}
	}

	return false
}

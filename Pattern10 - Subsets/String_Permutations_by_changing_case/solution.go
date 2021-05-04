package String_Permutations_by_changing_case

import (
	"strings"
	"unicode"
)

/*
Given a string, find all of its permutations preserving the character sequence but changing case.

Input: "ad52"
Output: "ad52", "Ad52", "aD52", "AD52"

Input: "ab7c"
Output: "ab7c", "Ab7c", "aB7c", "AB7c", "ab7C", "Ab7C", "aB7C", "AB7C"

ref: https://leetcode-cn.com/problems/letter-case-permutation/
*/

func letterCasePermutation(S string) []string {
	var ans []string
	ans = append(ans, "")

	for i := 0; i < len(S); i++ {
		if unicode.IsLetter(rune(S[i])) {
			n := len(ans)
			char := rune(S[i])
			for j := 0; j < n; j++ {
				tmp := ans[j]
				ans[j] = tmp + letterTo(char, false)
				ans = append(ans, tmp+letterTo(char, true))
			}
		} else {
			for j := 0; j < len(ans); j++ {
				ans[j] = ans[j] + string(S[i])
			}
		}
	}

	return ans
}

func letterTo(char rune, upper bool) string {
	charToString := string(char)
	if upper {
		return strings.ToUpper(charToString)
	} else {
		return strings.ToLower(charToString)
	}
}

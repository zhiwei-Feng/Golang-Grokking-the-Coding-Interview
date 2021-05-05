package Challenge1_Evaluate_Expression

import (
	"strconv"
	"unicode"
)

/*
Given an expression containing digits and operations (+, -, *),
find all possible ways in which the expression can be evaluated by grouping the numbers and operators using parentheses.

Input: "1+2*3"
Output: 7, 9
Explanation: 1+(2*3) => 7 and (1+2)*3 => 9

Input: "2*3-4-5"
Output: 8, -12, 7, -7, -3
Explanation: 2*(3-(4-5)) => 8, 2*(3-4-5) => -12, 2*3-(4-5) => 7, 2*(3-4)-5 => -7, (2*3)-4-5 => -3
*/

func diffWaysToEvaluateExpression(input string) []int {
	if v, err := strconv.Atoi(input); err == nil {
		return []int{v}
	}
	var ans []int
	for i := 0; i < len(input); i++ {
		chr := input[i]
		if unicode.IsDigit(rune(chr)) == false {
			leftPart := diffWaysToEvaluateExpression(input[:i])
			rightPart := diffWaysToEvaluateExpression(input[i+1:])

			for _, v1 := range leftPart {
				for _, v2 := range rightPart {
					switch chr {
					case '+':
						ans = append(ans, v1+v2)
					case '-':
						ans = append(ans, v1-v2)
					case '*':
						ans = append(ans, v1*v2)
					}
				}
			}
		}
	}

	return ans
}

package Balanced_Parentheses

/*
For a given number ‘N’, write a function to generate all combination of ‘N’ pairs of balanced parentheses.

Input: N=2
Output: (()), ()()

Input: N=3
Output: ((())), (()()), (())(), ()(()), ()()()

ref: https://leetcode-cn.com/problems/generate-parentheses/
*/

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var ans []string
	var queue []ParenthesesString
	queue = append(queue, ParenthesesString{"", 0, 0})
	for len(queue) != 0 {
		curStr := queue[0]
		queue = queue[1:] // pop
		if curStr.OpenCount == n && curStr.CloseCount == n {
			ans = append(ans, curStr.Value)
			continue
		}
		if curStr.OpenCount < n {
			queue = append(queue, ParenthesesString{curStr.Value + "(", curStr.OpenCount + 1, curStr.CloseCount})
		}

		if curStr.OpenCount > curStr.CloseCount {
			queue = append(queue, ParenthesesString{curStr.Value + ")", curStr.OpenCount, curStr.CloseCount + 1})
		}
	}

	return ans
}

type ParenthesesString struct {
	Value      string
	OpenCount  int // 多少个左括号
	CloseCount int // 多少个右括号
}

package Challenge2_Structurally_Unique_Binary_Search_Trees

/*
Given a number ‘n’, write a function to return all structurally unique Binary Search Trees (BST) that can store values 1 to ‘n’?

Input: 2
Output: List containing root nodes of all structurally unique BSTs.
Explanation: Here are the 2 structurally unique BSTs storing all numbers from 1 to 2:

Input: 3
Output: List containing root nodes of all structurally unique BSTs.
Explanation: Here are the 5 structurally unique BSTs storing all numbers from 1 to 3:

ref: https://leetcode-cn.com/problems/unique-binary-search-trees/
*/

func numTrees(n int) int {
	var bfs func(int, int) int
	var m = make(map[Range]int)
	bfs = func(start, end int) int {
		var res = 0
		if start == end {
			return 1
		}
		if v, ok := m[Range{start, end}]; ok {
			return v
		}
		for i := start; i <= end; i++ {
			if i == start {
				res += bfs(i+1, end)
			}
			if i == end {
				res += bfs(start, i-1)
			}
			left := bfs(start, i-1)
			right := bfs(i+1, end)
			res += left * right
		}
		m[Range{start, end}] = res
		return res
	}

	return bfs(1, n)
}

type Range struct {
	start int
	end   int
}

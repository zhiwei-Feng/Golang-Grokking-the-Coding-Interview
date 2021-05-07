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
ref: https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*TreeNode {
	var bfs func(int, int) []*TreeNode
	bfs = func(start, end int) []*TreeNode {
		if start == end {
			return []*TreeNode{{start, nil, nil}}
		}
		var res []*TreeNode
		for i := start; i <= end; i++ {
			if i == start {
				right := bfs(i+1, end)
				for _, v := range right {
					root := &TreeNode{i, nil, nil}
					root.Right = v
					res = append(res, root)
				}
				continue
			}
			if i == end {
				left := bfs(start, i-1)
				for _, v := range left {
					root := &TreeNode{i, nil, nil}
					root.Left = v
					res = append(res, root)
				}
				continue
			}
			left := bfs(start, i-1)
			right := bfs(i+1, end)
			for _, v1 := range left {
				for _, v2 := range right {
					root := &TreeNode{i, nil, nil}
					root.Left = v1
					root.Right = v2
					res = append(res, root)
				}
			}
		}

		return res
	}

	return bfs(1, n)
}

// https://leetcode-cn.com/problems/unique-binary-search-trees/
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
				continue
			}
			if i == end {
				res += bfs(start, i-1)
				continue
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

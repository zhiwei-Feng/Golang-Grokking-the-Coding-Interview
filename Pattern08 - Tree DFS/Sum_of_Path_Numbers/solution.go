package Sum_of_Path_Numbers

/*
Given a binary tree where each node can only have a digit (0-9) value, each root-to-leaf path will represent a number.
Find the total sum of all the numbers represented by all paths.

leetcode ref: https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	allSum := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, path int) {
		if root == nil {
			return
		}
		path = 10*path + root.Val
		if root.Left == nil && root.Right == nil {
			// leaf node
			allSum += path
		} else {
			if root.Left != nil {
				dfs(root.Left, path)
			}
			if root.Right != nil {
				dfs(root.Right, path)
			}
		}
	}
	dfs(root, 0)
	return allSum
}

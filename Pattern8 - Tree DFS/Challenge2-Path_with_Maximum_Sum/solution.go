package Challenge2_Path_with_Maximum_Sum

import "math"

/*
Find the path with the maximum sum in a given binary tree. Write a function that returns the maximum sum.

A path can be defined as a sequence of nodes between any two nodes and doesn’t necessarily pass through the root.
The path must contain at least one node.

ref: https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt8
	dfs(root, &maxSum)
	return maxSum
}

func dfs(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftMaxSum := max(dfs(node.Left, maxSum), 0)
	rightMaxSum := max(dfs(node.Right, maxSum), 0)

	currentSum := leftMaxSum + rightMaxSum + node.Val
	if currentSum > *maxSum {
		*maxSum = currentSum
	}

	return max(leftMaxSum, rightMaxSum) + node.Val
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

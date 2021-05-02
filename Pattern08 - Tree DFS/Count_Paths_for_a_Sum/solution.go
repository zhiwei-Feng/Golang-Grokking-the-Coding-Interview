package Count_Paths_for_a_Sum

/*
Given a binary tree and a number ‘S’, find all paths in the tree such that the sum of all the node values of each path equals ‘S’.
Please note that the paths can start or end at any node but all paths must follow direction from parent to child (top to bottom).

ref: https://leetcode-cn.com/problems/path-sum-iii/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var pathSumMap = make(map[int]int)
	pathSumMap[0] = 1
	return dfs(root, 0, pathSumMap, targetSum)
}

func dfs(node *TreeNode, acc int, sumMap map[int]int, target int) int {
	if node == nil {
		return 0
	}
	var res int
	acc += node.Val
	res += sumMap[acc-target]
	sumMap[acc]++

	res += dfs(node.Left, acc, sumMap, target)
	res += dfs(node.Right, acc, sumMap, target)

	sumMap[acc]--

	return res
}

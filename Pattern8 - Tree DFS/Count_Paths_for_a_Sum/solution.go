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
	if root == nil {
		return 0
	}
	sum := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, left int) {
		if root.Val == left {
			sum++
		}
		if root.Left != nil {
			dfs(root.Left, left-root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, left-root.Val)
		}
	}

	// bfs+dfs
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			dfs(curNode, targetSum)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}
	return sum
}

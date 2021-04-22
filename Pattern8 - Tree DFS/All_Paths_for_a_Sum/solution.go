package All_Paths_for_a_Sum

/*
Given a binary tree and a number ‘S’, find all paths from root-to-leaf such that the sum of all the node values of each path equals ‘S’.

leetcode ref: https://leetcode-cn.com/problems/path-sum-ii/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, left int) {
		if root == nil {
			return
		}
		left -= root.Val
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()
		if left == 0 && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(root.Left, left)
		dfs(root.Right, left)
	}
	dfs(root, targetSum)
	return
}

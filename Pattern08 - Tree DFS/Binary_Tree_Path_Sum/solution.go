package Binary_Tree_Path_Sum

/*
Given a binary tree and a number ‘S’, find if the tree has a path from root-to-leaf such that the sum of all the node values of that path equals ‘S’.
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 因为每一次向下遍历，targetSum都会减去遍历后的节点值，因此如果存在一条这样的path，那么遍历到该path的叶子节点时
	// 其targetSum一定会等于叶子节点的值
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

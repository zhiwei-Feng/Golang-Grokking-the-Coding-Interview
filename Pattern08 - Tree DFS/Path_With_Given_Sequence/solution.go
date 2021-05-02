package Path_With_Given_Sequence

/*
Given a binary tree and a number sequence, find if the sequence is present as a root-to-leaf path in the given tree.

ref: https://leetcode-cn.com/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil {
		return 0 == len(arr)
	}
	return dfs(root, 0, arr)
}

func dfs(root *TreeNode, arrInd int, arr []int) bool {
	if root == nil {
		// 说明该path虽然前面都匹配，但长度不够
		return false
	}
	if arrInd >= len(arr) || root.Val != arr[arrInd] {
		// path前面匹配但是path过长，或者当前路径不匹配
		return false
	}
	if arrInd == len(arr)-1 && root.Left == nil && root.Right == nil {
		// 必须加上arrInd == len(arr)来保证完全匹配
		return true
	}
	return dfs(root.Left, arrInd+1, arr) || dfs(root.Right, arrInd+1, arr)
}

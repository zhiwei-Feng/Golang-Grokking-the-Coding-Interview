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
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, left int) {
		if root == nil {
			return
		}
		left -= root.Val
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()
		if left == 0 && root.Left == nil && root.Right == nil {
			// 解释下为什么要这么做，而不是
			// ans = append(ans, path)
			// 因为如果像上面做，其实在ans存放的是path指针所指向的地址区，因此path的变化会影响到ans的内容
			// 而我们想要的是当ans append操作的时候，是加入一个immutable的[]int，
			// 因此必须使用append([]int{}, path...)重新分配一块新的内存空间存放path的复制体，并用一个新的指针指向
			// 而此时ans保存的是新的指针指向的地址区，而且不会在随后的代码执行中被修改。
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(root.Left, left)
		dfs(root.Right, left)
	}
	dfs(root, targetSum)
	return
}

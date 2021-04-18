package Minimum_Depth_of_a_Binary_Tree

/*
Find the minimum depth of a binary tree.
The minimum depth is the number of nodes along the shortest path from the root node to the nearest leaf node.
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue    = make([]*TreeNode, 0, 10)
		minDepth = 0
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		minDepth++
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left == nil && curNode.Right == nil {
				return minDepth
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}

	return minDepth
}

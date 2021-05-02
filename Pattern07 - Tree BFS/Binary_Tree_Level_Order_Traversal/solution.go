package Binary_Tree_Level_Order_Traversal

/*
Given a binary tree, populate an array to represent its level-by-level traversal.
You should populate the values of all nodes of each level from left to right in separate sub-arrays.

Leetcode ref: problems/binary-tree-level-order-traversal/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(head *TreeNode) [][]int {
	if head == nil {
		return [][]int{}
	}
	var (
		result = make([][]int, 0, 10)
		queue  = make([]*TreeNode, 0, 10)
	)

	queue = append(queue, head)
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevelArray := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			currentLevelArray = append(currentLevelArray, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		result = append(result, currentLevelArray)
	}

	return result
}

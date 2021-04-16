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
		result              = make([][]int, 0, 10)
		queue               = make([]*TreeNode, 0, 10)
		lastNodeOfPrevLevel *TreeNode
		currentLevel        = 0
		tmpLastNode         *TreeNode
	)

	queue = append(queue, head)
	lastNodeOfPrevLevel = head
	result = append(result, make([]int, 0, 10))
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		result[currentLevel] = append(result[currentLevel], currentNode.Val)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
			tmpLastNode = currentNode.Left
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
			tmpLastNode = currentNode.Right
		}

		if currentNode == lastNodeOfPrevLevel && len(queue) > 0 {
			lastNodeOfPrevLevel = tmpLastNode
			result = append(result, make([]int, 0, 10))
			currentLevel++
		}
	}

	return result
}

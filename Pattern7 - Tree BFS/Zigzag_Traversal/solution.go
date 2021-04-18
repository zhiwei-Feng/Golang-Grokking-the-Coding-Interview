package Zigzag_Traversal

/*
Given a binary tree, populate an array to represent its zigzag level order traversal.
You should populate the values of all nodes of the first level from left to right,
then right to left for the next level and keep alternating in the same manner for the following levels.
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		result      = make([][]int, 0)
		queue       = make([]*TreeNode, 0)
		leftToRight = true
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		levelList := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}

			if leftToRight {
				levelList[i] = curNode.Val
			} else {
				levelList[levelSize-1-i] = curNode.Val
			}
		}
		leftToRight = !leftToRight
		result = append(result, levelList)
	}

	return result
}

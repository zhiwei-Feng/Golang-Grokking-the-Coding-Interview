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

func reverse(arr []int) {
	var (
		i = 0
		j = len(arr) - 1
	)
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		result     = make([][]int, 0)
		queue      = make([]*TreeNode, 0)
		zigReverse = false
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		levelList := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			levelList = append(levelList, curNode.Val)
		}
		if zigReverse {
			reverse(levelList)
		}
		zigReverse = !zigReverse
		result = append(result, levelList)
	}

	return result
}

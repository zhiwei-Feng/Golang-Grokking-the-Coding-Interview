package Reverse_Level_Order_Traversal

/*
Given a binary tree, populate an array to represent its level-by-level traversal in reverse order, i.e., the lowest level comes first.
You should populate the values of all nodes in each level from left to right in separate sub-arrays.

Leetcode ref: problems/binary-tree-level-order-traversal-ii
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		result = make([][]int, 0, 10)
		queue  = make([]*TreeNode, 0, 10)
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevelList := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			currentLevelList = append(currentLevelList, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		// 下面这种写法会有更大的内存占用
		//lefJoin := [][]int{currentLevelList}
		//result = append(lefJoin, result...)
		result = append(result, currentLevelList)
	}

	// reverse result list
	tmp := result
	result = make([][]int, 0, 10)
	for i := len(tmp) - 1; i >= 0; i-- {
		result = append(result, tmp[i])
	}

	return result
}

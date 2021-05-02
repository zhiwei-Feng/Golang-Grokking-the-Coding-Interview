package Level_Order_Successor

/*
Given a binary tree and a node, find the level order successor of the given node in the tree.
The level order successor is the node that appears right after the given node in the level order traversal.
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findSuccessor(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var queue = make([]*TreeNode, 0)
	var hit = false
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if hit {
				return curNode
			}
			if curNode.Val == key {
				hit = true
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}

	return nil
}

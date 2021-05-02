package Connect_Level_Order_Siblings

/*
Given a binary tree, connect each node with its level order successor.
The last node of each level should point to a null node.
*/

type Node struct {
	Val         int
	Left, Right *Node
	Next        *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		var previousNode *Node
		//p := make([]*Node, 0)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if previousNode != nil {
				previousNode.Next = curNode
			}
			previousNode = curNode
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}
	return root
}

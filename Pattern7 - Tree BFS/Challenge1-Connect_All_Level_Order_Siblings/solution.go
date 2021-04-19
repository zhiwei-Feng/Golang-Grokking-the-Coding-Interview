package Challenge1_Connect_All_Level_Order_Siblings

/*
Given a binary tree, connect each node with its level order successor.
The last node of each level should point to the first node of the next level.
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var (
		queue    = make([]*Node, 0)
		prevNode *Node
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if prevNode != nil {
				prevNode.Next = curNode
			}
			prevNode = curNode
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

package Challenge2_Right_View_of_a_Binary_Tree

/*
Given a binary tree, return an array containing nodes in its right view.
The right view of a binary tree is the set of nodes visible when the tree is seen from the right side.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		rightViewList = make([]int, 0)
		queue         = make([]*TreeNode, 0)
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				rightViewList = append(rightViewList, curNode.Val)
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}

	return rightViewList
}

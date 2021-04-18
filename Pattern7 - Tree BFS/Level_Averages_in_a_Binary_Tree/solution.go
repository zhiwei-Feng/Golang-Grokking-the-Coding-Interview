package Level_Averages_in_a_Binary_Tree

/*
Given a binary tree, populate an array to represent the averages of all of its levels.
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var (
		result = make([]float64, 0, 10)
		queue  = make([]*TreeNode, 0, 10)
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0.0
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			levelSum += float64(curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		result = append(result, levelSum/float64(levelSize))
	}

	return result
}

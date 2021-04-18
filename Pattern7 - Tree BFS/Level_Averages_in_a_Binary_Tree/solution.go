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
		p := make([]*TreeNode, 0, levelSize*2)
		for i := 0; i < levelSize; i++ {
			curNode := queue[i]
			levelSum += float64(curNode.Val)
			if curNode.Left != nil {
				p = append(p, curNode.Left)
			}
			if curNode.Right != nil {
				p = append(p, curNode.Right)
			}
		}
		queue = p
		result = append(result, levelSum/float64(levelSize))
	}

	return result
}
